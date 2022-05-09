// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nsxtreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver"

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/model/pdata"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver/internal/metadata"
	dm "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver/internal/model"
)

type scraper struct {
	config   *Config
	settings component.TelemetrySettings
	host     component.Host
	client   Client
	mb       *metadata.MetricsBuilder
	logger   *zap.Logger
}

func newScraper(cfg *Config, settings component.TelemetrySettings) *scraper {
	return &scraper{
		config:   cfg,
		settings: settings,
		mb:       metadata.NewMetricsBuilder(cfg.Metrics),
		logger:   settings.Logger,
	}
}

func (s *scraper) start(ctx context.Context, host component.Host) error {
	s.host = host
	err := s.ensureClient()
	// allow reconnectivity
	if err != nil {
		return fmt.Errorf("unable to ensure client connectivity: %s", err)
	}
	return nil
}

type nodeClass int

const (
	transportClass nodeClass = iota
	managerClass
)

func (s *scraper) scrape(ctx context.Context) (pdata.Metrics, error) {
	errs := &scrapererror.ScrapeErrors{}
	if err := s.ensureClient(); err != nil {
		errs.Add(err)
		return pmetric.NewMetrics(), errs.Combine()
	}

	r := s.retrieve(ctx, errs)

	colTime := pdata.NewTimestampFromTime(time.Now())
	s.process(r, colTime)

	return s.mb.Emit(), errs.Combine()
}

type nodeInfo struct {
	nodeProps  dm.NodeProperties
	nodeType   string
	interfaces []interfaceInformation
	stats      *dm.NodeStatus
}

type interfaceInformation struct {
	iFace dm.NetworkInterface
	stats *dm.NetworkInterfaceStats
}

func (s *scraper) retrieve(ctx context.Context, errs *scrapererror.ScrapeErrors) []*nodeInfo {
	r := []*nodeInfo{}

	tNodes, err := s.client.TransportNodes(ctx)
	if err != nil {
		errs.AddPartial(1, err)
		return r
	}

	cNodes, err := s.client.ClusterNodes(ctx)
	if err != nil {
		errs.AddPartial(1, err)
		return r
	}

	wg := &sync.WaitGroup{}
	for _, n := range tNodes {
		nodeInfo := &nodeInfo{
			nodeProps: n.NodeProperties,
			nodeType:  "transport",
		}
		wg.Add(2)
		go s.retrieveInterfaces(ctx, n.NodeProperties, nodeInfo, transportClass, wg, errs)
		go s.retrieveNodeStats(ctx, n.NodeProperties, nodeInfo, transportClass, wg, errs)

		r = append(r, nodeInfo)
	}

	for _, n := range cNodes {
		// no useful stats are recorded for controller nodes
		if clusterNodeType(n) != "manager" {
			continue
		}

		nodeInfo := &nodeInfo{
			nodeProps: n.NodeProperties,
			nodeType:  "manager",
		}

		wg.Add(2)
		go s.retrieveInterfaces(ctx, n.NodeProperties, nodeInfo, managerClass, wg, errs)
		go s.retrieveNodeStats(ctx, n.NodeProperties, nodeInfo, managerClass, wg, errs)

		r = append(r, nodeInfo)
	}

	wg.Wait()

	return r
}

func (s *scraper) retrieveInterfaces(
	ctx context.Context,
	nodeProps dm.NodeProperties,
	nodeInfo *nodeInfo,
	nodeClass nodeClass,
	wg *sync.WaitGroup,
	errs *scrapererror.ScrapeErrors,
) {
	defer wg.Done()
	interfaces, err := s.client.Interfaces(ctx, nodeProps.ID, nodeClass)
	if err != nil {
		errs.AddPartial(1, err)
		return
	}
	nodeInfo.interfaces = []interfaceInformation{}
	for _, i := range interfaces {
		interfaceInfo := interfaceInformation{
			iFace: i,
		}
		stats, err := s.client.InterfaceStatus(ctx, nodeProps.ID, i.InterfaceId, nodeClass)
		if err != nil {
			errs.AddPartial(1, err)
		}
		interfaceInfo.stats = stats
		nodeInfo.interfaces = append(nodeInfo.interfaces, interfaceInfo)
	}
}

func (s *scraper) retrieveNodeStats(
	ctx context.Context,
	nodeProps dm.NodeProperties,
	nodeInfo *nodeInfo,
	nodeClass nodeClass,
	wg *sync.WaitGroup,
	errs *scrapererror.ScrapeErrors,
) {
	defer wg.Done()
	ns, err := s.client.NodeStatus(ctx, nodeProps.ID, nodeClass)
	if err != nil {
		errs.AddPartial(1, err)
		return
	}
	nodeInfo.stats = ns
}

func (s *scraper) process(
	nodes []*nodeInfo,
	colTime pdata.Timestamp,
) {
	for _, n := range nodes {
		for _, i := range n.interfaces {
			s.recordNodeInterface(colTime, n.nodeProps, i)
		}
		s.recordNode(colTime, n)
	}
}

func (s *scraper) recordNodeInterface(colTime pdata.Timestamp, nodeProps dm.NodeProperties, i interfaceInformation) {
	s.mb.RecordNsxInterfacePacketCountDataPoint(colTime, i.stats.RxDropped, metadata.AttributeDirectionReceived, metadata.AttributePacketTypeDropped)
	s.mb.RecordNsxInterfacePacketCountDataPoint(colTime, i.stats.RxErrors, metadata.AttributeDirectionReceived, metadata.AttributePacketTypeErrored)
	successRxPackets := i.stats.RxPackets - i.stats.RxDropped - i.stats.RxErrors
	s.mb.RecordNsxInterfacePacketCountDataPoint(colTime, successRxPackets, metadata.AttributeDirectionReceived, metadata.AttributePacketTypeSuccess)

	s.mb.RecordNsxInterfacePacketCountDataPoint(colTime, i.stats.TxDropped, metadata.AttributeDirectionTransmitted, metadata.AttributePacketTypeDropped)
	s.mb.RecordNsxInterfacePacketCountDataPoint(colTime, i.stats.TxErrors, metadata.AttributeDirectionTransmitted, metadata.AttributePacketTypeErrored)
	successTxPackets := i.stats.TxPackets - i.stats.TxDropped - i.stats.TxErrors
	s.mb.RecordNsxInterfacePacketCountDataPoint(colTime, successTxPackets, metadata.AttributeDirectionTransmitted, metadata.AttributePacketTypeSuccess)

	s.mb.RecordNsxInterfaceThroughputDataPoint(colTime, i.stats.RxBytes, metadata.AttributeDirectionReceived)
	s.mb.RecordNsxInterfaceThroughputDataPoint(colTime, i.stats.TxBytes, metadata.AttributeDirectionTransmitted)

	s.mb.EmitForResource(
		metadata.WithNsxInterfaceID(i.iFace.InterfaceId),
		metadata.WithNsxNodeName(nodeProps.Name),
	)
}

func (s *scraper) recordNode(
	colTime pdata.Timestamp,
	info *nodeInfo,
) {
	if info.stats == nil {
		return
	}

	ss := info.stats.SystemStatus
	s.mb.RecordNsxNodeCPUUtilizationDataPoint(colTime, ss.CPUUsage.AvgCPUCoreUsageDpdk, metadata.AttributeCPUProcessClassDatapath)
	s.mb.RecordNsxNodeCPUUtilizationDataPoint(colTime, ss.CPUUsage.AvgCPUCoreUsageNonDpdk, metadata.AttributeCPUProcessClassServices)
	s.mb.RecordNsxNodeMemoryUsageDataPoint(colTime, int64(ss.MemUsed))

	if ss.EdgeMemUsage != nil {
		s.mb.RecordNsxNodeCacheMemoryUsageDataPoint(colTime, int64(ss.EdgeMemUsage.CacheUsage))
	}

	s.mb.RecordNsxNodeDiskUsageDataPoint(colTime, int64(ss.DiskSpaceUsed), metadata.AttributeDiskStateUsed)
	availableStorage := ss.DiskSpaceTotal - ss.DiskSpaceUsed
	s.mb.RecordNsxNodeDiskUsageDataPoint(colTime, int64(availableStorage), metadata.AttributeDiskStateAvailable)
	// ensure division by zero is safeguarded
	s.mb.RecordNsxNodeDiskUtilizationDataPoint(colTime, float64(ss.DiskSpaceUsed)/float64(math.Max(float64(ss.DiskSpaceTotal), 1)))

	s.mb.EmitForResource(
		metadata.WithNsxNodeName(info.nodeProps.Name),
		metadata.WithNsxNodeID(info.nodeProps.ID),
		metadata.WithNsxNodeType(info.nodeType),
	)
}

func (s *scraper) ensureClient() error {
	if s.client != nil {
		return nil
	}
	client, err := newClient(s.config, s.settings, s.host, s.logger.Named("client"))
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

func clusterNodeType(node dm.ClusterNode) string {
	if node.ControllerRole != nil {
		return "controller"
	}
	return "manager"
}
