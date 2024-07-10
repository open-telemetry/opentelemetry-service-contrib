// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	vt "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan"
	"github.com/vmware/govmomi/vsan/types"
	"go.uber.org/zap"
)

// vcenterClient is a client that collects data from a vCenter endpoint.
type vcenterClient struct {
	logger     *zap.Logger
	moClient   *govmomi.Client
	vimDriver  *vim25.Client
	vsanDriver *vsan.Client
	finder     *find.Finder
	pm         *performance.Manager
	vm         *view.Manager
	cfg        *Config
}

var newVcenterClient = defaultNewVcenterClient

func defaultNewVcenterClient(l *zap.Logger, c *Config) *vcenterClient {
	return &vcenterClient{
		logger: l,
		cfg:    c,
	}
}

// EnsureConnection will establish a connection to the vSphere SDK if not already established
func (vc *vcenterClient) EnsureConnection(ctx context.Context) error {
	if vc.moClient != nil {
		sessionActive, _ := vc.moClient.SessionManager.SessionIsActive(ctx)
		if sessionActive {
			return nil
		}
	}

	sdkURL, err := vc.cfg.SDKUrl()
	if err != nil {
		return err
	}
	client, err := govmomi.NewClient(ctx, sdkURL, vc.cfg.Insecure)
	if err != nil {
		return fmt.Errorf("unable to connect to vSphere SDK on listed endpoint: %w", err)
	}
	tlsCfg, err := vc.cfg.LoadTLSConfig(ctx)
	if err != nil {
		return err
	}
	if tlsCfg != nil {
		client.DefaultTransport().TLSClientConfig = tlsCfg
	}
	user := url.UserPassword(vc.cfg.Username, string(vc.cfg.Password))
	err = client.Login(ctx, user)
	if err != nil {
		return fmt.Errorf("unable to login to vcenter sdk: %w", err)
	}
	vc.moClient = client
	vc.vimDriver = client.Client
	vc.finder = find.NewFinder(vc.vimDriver)
	vc.pm = performance.NewManager(vc.vimDriver)
	vc.vm = view.NewManager(vc.vimDriver)
	vsanDriver, err := vsan.NewClient(ctx, vc.vimDriver)
	if err != nil {
		vc.logger.Info(fmt.Errorf("could not create VSAN client: %w", err).Error())
	} else {
		vc.vsanDriver = vsanDriver
	}
	return nil
}

// Disconnect will logout of the autenticated session
func (vc *vcenterClient) Disconnect(ctx context.Context) error {
	if vc.moClient != nil {
		return vc.moClient.Logout(ctx)
	}
	return nil
}

// Datacenters returns the Datacenters of the vSphere SDK
func (vc *vcenterClient) Datacenters(ctx context.Context) ([]mo.Datacenter, error) {
	v, err := vc.vm.CreateContainerView(ctx, vc.vimDriver.ServiceContent.RootFolder, []string{"Datacenter"}, true)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Datacenters: %w", err)
	}
	defer func() { _ = v.Destroy(ctx) }()

	var datacenters []mo.Datacenter
	err = v.Retrieve(ctx, []string{"Datacenter"}, []string{
		"name",
	}, &datacenters)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Datacenters: %w", err)
	}

	return datacenters, nil
}

// Datastores returns the Datastores of the vSphere SDK
func (vc *vcenterClient) Datastores(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.Datastore, error) {
	v, err := vc.vm.CreateContainerView(ctx, containerMoRef, []string{"Datastore"}, true)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Datastores: %w", err)
	}
	defer func() { _ = v.Destroy(ctx) }()

	var datastores []mo.Datastore
	err = v.Retrieve(ctx, []string{"Datastore"}, []string{
		"name",
		"summary.capacity",
		"summary.freeSpace",
	}, &datastores)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Datastores: %w", err)
	}

	return datastores, nil
}

// ComputeResources returns the ComputeResources (& ClusterComputeResources) of the vSphere SDK
func (vc *vcenterClient) ComputeResources(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.ComputeResource, error) {
	v, err := vc.vm.CreateContainerView(ctx, containerMoRef, []string{"ComputeResource"}, true)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve ComputeResources (& ClusterComputeResources): %w", err)
	}
	defer func() { _ = v.Destroy(ctx) }()

	var computes []mo.ComputeResource
	err = v.Retrieve(ctx, []string{"ComputeResource"}, []string{
		"name",
		"datastore",
		"host",
		"summary",
	}, &computes)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve ComputeResources (& ClusterComputeResources): %w", err)
	}

	return computes, nil
}

// HostSystems returns the HostSystems of the vSphere SDK
func (vc *vcenterClient) HostSystems(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.HostSystem, error) {
	v, err := vc.vm.CreateContainerView(ctx, containerMoRef, []string{"HostSystem"}, true)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve HostSystems: %w", err)
	}
	defer func() { _ = v.Destroy(ctx) }()

	var hosts []mo.HostSystem
	err = v.Retrieve(ctx, []string{"HostSystem"}, []string{
		"name",
		"runtime.powerState",
		"summary.hardware.memorySize",
		"summary.hardware.numCpuCores",
		"summary.hardware.cpuMhz",
		"summary.quickStats.overallMemoryUsage",
		"summary.quickStats.overallCpuUsage",
		"summary.overallStatus",
		"vm",
		"parent",
	}, &hosts)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve HostSystems: %w", err)
	}

	return hosts, nil
}

// ResourcePools returns the ResourcePools (&VirtualApps) of the vSphere SDK
func (vc *vcenterClient) ResourcePools(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.ResourcePool, error) {
	v, err := vc.vm.CreateContainerView(ctx, containerMoRef, []string{"ResourcePool"}, true)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve ResourcePools (&VirtualApps): %w", err)
	}
	defer func() { _ = v.Destroy(ctx) }()

	var rps []mo.ResourcePool
	err = v.Retrieve(ctx, []string{"ResourcePool"}, []string{
		"summary",
		"name",
		"owner",
		"vm",
	}, &rps)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve ResourcePools (&VirtualApps): %w", err)
	}

	return rps, nil
}

// VMS returns the VirtualMachines of the vSphere SDK
func (vc *vcenterClient) VMs(ctx context.Context, containerMoRef vt.ManagedObjectReference) ([]mo.VirtualMachine, error) {
	v, err := vc.vm.CreateContainerView(ctx, containerMoRef, []string{"VirtualMachine"}, true)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve VMs: %w", err)
	}
	defer func() { _ = v.Destroy(ctx) }()

	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{
		"name",
		"config.hardware.numCPU",
		"config.instanceUuid",
		"config.template",
		"runtime.powerState",
		"runtime.maxCpuUsage",
		"summary.quickStats.guestMemoryUsage",
		"summary.quickStats.balloonedMemory",
		"summary.quickStats.swappedMemory",
		"summary.quickStats.ssdSwappedMemory",
		"summary.quickStats.overallCpuUsage",
		"summary.overallStatus",
		"summary.config.memorySizeMB",
		"summary.storage.committed",
		"summary.storage.uncommitted",
		"summary.runtime.host",
		"resourcePool",
		"parentVApp",
	}, &vms)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve VMs: %w", err)
	}

	return vms, nil
}

// DatacenterInventoryListObjects returns the Datacenters (with populated InventoryLists) of the vSphere SDK
func (vc *vcenterClient) DatacenterInventoryListObjects(ctx context.Context) ([]*object.Datacenter, error) {
	dcs, err := vc.finder.DatacenterList(ctx, "*")
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Datacenters with InventoryLists: %w", err)
	}

	return dcs, nil
}

// ResourcePoolInventoryListObjects returns the ResourcePools (with populated InventoryLists) of the vSphere SDK
func (vc *vcenterClient) ResourcePoolInventoryListObjects(
	ctx context.Context,
	dcs []*object.Datacenter,
) ([]*object.ResourcePool, error) {
	allRPools := []*object.ResourcePool{}
	for _, dc := range dcs {
		vc.finder.SetDatacenter(dc)
		rps, err := vc.finder.ResourcePoolList(ctx, "*")
		var notFoundErr *find.NotFoundError
		if err != nil && !errors.As(err, &notFoundErr) {
			return nil, fmt.Errorf("unable to retrieve ResourcePools with InventoryLists for datacenter %s: %w", dc.InventoryPath, err)
		}
		allRPools = append(allRPools, rps...)
	}

	return allRPools, nil
}

// VAppInventoryListObjects returns the vApps (with populated InventoryLists) of the vSphere SDK
func (vc *vcenterClient) VAppInventoryListObjects(
	ctx context.Context,
	dcs []*object.Datacenter,
) ([]*object.VirtualApp, error) {
	allVApps := []*object.VirtualApp{}
	for _, dc := range dcs {
		vc.finder.SetDatacenter(dc)
		vApps, err := vc.finder.VirtualAppList(ctx, "*")
		if err == nil {
			allVApps = append(allVApps, vApps...)
			continue
		}

		var notFoundErr *find.NotFoundError
		if !errors.As(err, &notFoundErr) {
			return nil, fmt.Errorf("unable to retrieve vApps with InventoryLists for datacenter %s: %w", dc.InventoryPath, err)
		}
	}

	return allVApps, nil
}

// PerfMetricsQueryResult contains performance metric related data
type PerfMetricsQueryResult struct {
	// Contains performance metrics keyed by MoRef string
	resultsByRef map[string]*performance.EntityMetric
}

// PerfMetricsQuery returns the requested performance metrics for the requested resources
// over a given sample interval and sample count
func (vc *vcenterClient) PerfMetricsQuery(
	ctx context.Context,
	spec vt.PerfQuerySpec,
	names []string,
	objs []vt.ManagedObjectReference,
) (*PerfMetricsQueryResult, error) {
	if vc.pm == nil {
		return &PerfMetricsQueryResult{}, nil
	}
	vc.pm.Sort = true
	sample, err := vc.pm.SampleByName(ctx, spec, names, objs)
	if err != nil {
		return nil, err
	}
	result, err := vc.pm.ToMetricSeries(ctx, sample)
	if err != nil {
		return nil, err
	}

	resultsByRef := map[string]*performance.EntityMetric{}
	for i := range result {
		resultsByRef[result[i].Entity.Value] = &result[i]
	}
	return &PerfMetricsQueryResult{
		resultsByRef: resultsByRef,
	}, nil
}

// VSANQueryResults contains all returned vSAN metric related data
type VSANQueryResults struct {
	// Contains vSAN metric data keyed by UUID string
	MetricResultsByUUID map[string]*VSANMetricResults
}

// VSANMetricResults contains vSAN metric related data for a single resource
type VSANMetricResults struct {
	// Contains UUID info for related resource
	UUID string
	// Contains returned metric value info for all metrics
	MetricDetails []*VSANMetricDetails
}

// VSANMetricDetails contains vSAN metric data for a single metric
type VSANMetricDetails struct {
	// Contains the metric label
	MetricLabel string
	// Contains vSAN metric values keyed by timestamp
	ValuesByTimestamp map[time.Time]int64
}

// VSANVirtualMachines returns back virtual machine vSAN performance metrics for a group of clusters
func (vc *vcenterClient) VSANVirtualMachines(
	ctx context.Context,
	clusterRefs []*vt.ManagedObjectReference,
) (*VSANQueryResults, error) {
	allResults := VSANQueryResults{
		MetricResultsByUUID: map[string]*VSANMetricResults{},
	}

	for _, clusterRef := range clusterRefs {
		results, err := vc.vSANVirtualMachinesByCluster(ctx, clusterRef)
		// Check for expected errors
		if err != nil {
			faultErr := errors.Unwrap(err)
			if faultErr == nil {
				return nil, fmt.Errorf("problem retrieving Virtual Machine vSAN metrics: %w", err)
			}
			if !soap.IsSoapFault(faultErr) {
				return nil, fmt.Errorf("problem retrieving Virtual Machine vSAN metrics: %w", err)
			}

			fault := soap.ToSoapFault(faultErr)
			msg := fault.String
			if fault.Detail.Fault != nil {
				msg = reflect.TypeOf(fault.Detail.Fault).Name()
			}
			switch msg {
			case "NotSupported":
				vc.logger.Info(fmt.Sprintf("problem retrieving Virtual Machine vSAN metrics: %s", err.Error()))
				return &allResults, nil
			case "NotFound":
				vc.logger.Debug(fmt.Sprintf("no Virtual Machine vSAN metrics found: %s", err.Error()))
				return &allResults, nil
			default:
				return nil, fmt.Errorf("problem retrieving Virtual Machine vSAN metrics: %w", err)
			}
		}

		maps.Copy(allResults.MetricResultsByUUID, results.MetricResultsByUUID)
	}

	return &allResults, nil
}

// vSANVirtualMachinesByCluster returns back virtual machine vSAN performance metrics for a cluster
func (vc *vcenterClient) vSANVirtualMachinesByCluster(
	ctx context.Context,
	clusterRef *vt.ManagedObjectReference,
) (*VSANQueryResults, error) {
	queryResults := VSANQueryResults{
		MetricResultsByUUID: map[string]*VSANMetricResults{},
	}
	// Not all vCenters support vSAN so just return an empty result
	if vc.vsanDriver == nil {
		return &queryResults, nil
	}

	now := time.Now()
	querySpec := []types.VsanPerfQuerySpec{
		{
			EntityRefId: "virtual-machine:*",
			StartTime:   &now,
			EndTime:     &now,
			Labels: []string{
				"iopsRead",
				"iopsWrite",
				"throughputRead",
				"throughputWrite",
				"latencyRead",
				"latencyWrite",
			},
		},
	}
	rawResults, err := vc.vsanDriver.VsanPerfQueryPerf(ctx, clusterRef, querySpec)
	if err != nil {
		return nil, fmt.Errorf("problem retrieving Virtual Machine vSAN metrics for cluster %s: %w", clusterRef.Value, err)
	}

	queryResults.MetricResultsByUUID = map[string]*VSANMetricResults{}
	for _, rawResult := range rawResults {
		metricResults, err := convertVSANResultToMetricResults(rawResult)
		if err != nil && metricResults != nil {
			return &queryResults, fmt.Errorf("problem retrieving Virtual Machine [%s] vSAN metrics for cluster %s: %w", metricResults.UUID, clusterRef.Value, err)
		}
		if err != nil {
			return &queryResults, fmt.Errorf("problem retrieving Virtual Machine vSAN metrics for cluster %s: %w", clusterRef.Value, err)
		}

		queryResults.MetricResultsByUUID[metricResults.UUID] = metricResults
	}

	return &queryResults, nil
}

func convertVSANResultToMetricResults(vSANResult types.VsanPerfEntityMetricCSV) (*VSANMetricResults, error) {
	uuid, err := uuidFromEntityRefID(vSANResult.EntityRefId)
	if err != nil {
		return nil, err
	}

	metricResults := VSANMetricResults{
		UUID:          uuid,
		MetricDetails: []*VSANMetricDetails{},
	}

	// Parse all timestamps
	timeStrings := strings.Split(vSANResult.SampleInfo, ",")
	timestamps := []time.Time{}
	for _, timeString := range timeStrings {
		timestamp, err := time.Parse("2006-01-02 15:04:05", timeString)
		if err != nil {
			return &metricResults, fmt.Errorf("problem parsing timestamp from %s: %w", timeString, err)
		}

		timestamps = append(timestamps, timestamp)
	}

	// Parse all metrics
	for _, vSANValue := range vSANResult.Value {
		metricDetails, err := convertVSANValueToMetricDetails(vSANValue, timestamps)
		if err != nil {
			return &metricResults, err
		}

		metricResults.MetricDetails = append(metricResults.MetricDetails, metricDetails)
	}
	return &metricResults, nil
}

func convertVSANValueToMetricDetails(vSANValue types.VsanPerfMetricSeriesCSV, timestamps []time.Time) (*VSANMetricDetails, error) {
	metricLabel := vSANValue.MetricId.Label
	metricDetails := VSANMetricDetails{
		MetricLabel:       metricLabel,
		ValuesByTimestamp: map[time.Time]int64{},
	}
	valueStrings := strings.Split(vSANValue.Values, ",")
	if len(valueStrings) != len(timestamps) {
		return nil, fmt.Errorf("number of timestamps [%d] doesn't match number of values [%d] for metric %s", len(timestamps), len(valueStrings), metricLabel)
	}

	// Match up timestamps with metric values
	for i, valueString := range valueStrings {
		value, err := strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("problem converting value [%s] for metric %s", valueString, metricLabel)
		}

		metricDetails.ValuesByTimestamp[timestamps[i]] = value
	}

	return &metricDetails, nil
}

// uuidFromEntityRefID returns the UUID portion of the EntityRefId
func uuidFromEntityRefID(id string) (string, error) {
	colonIndex := strings.Index(id, ":")
	if colonIndex != -1 {
		uuid := id[colonIndex+1:]
		return uuid, nil
	}

	return "", fmt.Errorf("no ':' found in EntityRefId [%s] to parse UUID", id)
}
