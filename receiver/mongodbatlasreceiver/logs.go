// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context"
	"errors"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	rcvr "go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/model"
)

const mongoDBMajorVersion4_2 = "4.2"

type logsReceiver struct {
	log         *zap.Logger
	cfg         *Config
	client      *internal.MongoDBAtlasClient
	consumer    consumer.Logs
	stopperChan chan struct{}
	wg          sync.WaitGroup
	start       time.Time
	end         time.Time
}

type ProjectContext struct {
	Project mongodbatlas.Project
	orgName string
}

// MongoDB Atlas Documentation reccommends a polling interval of 5  minutes: https://www.mongodb.com/docs/atlas/reference/api/logs/#logs
const collectionInterval = time.Minute * 5

func newMongoDBAtlasLogsReceiver(settings rcvr.CreateSettings, cfg *Config, consumer consumer.Logs) *logsReceiver {
	client := internal.NewMongoDBAtlasClient(cfg.PublicKey, string(cfg.PrivateKey), cfg.RetrySettings, settings.Logger)

	for _, p := range cfg.Logs.Projects {
		p.populateIncludesAndExcludes()
	}

	return &logsReceiver{
		log:         settings.Logger,
		cfg:         cfg,
		client:      client,
		stopperChan: make(chan struct{}),
		consumer:    consumer,
	}
}

// Log receiver logic
func (s *logsReceiver) Start(ctx context.Context, _ component.Host) error {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.start = time.Now().Add(-collectionInterval)
		s.end = time.Now()
		for {
			s.collect(ctx)
			// collection interval loop,
			select {
			case <-ctx.Done():
				return
			case <-s.stopperChan:
				return
			case <-time.After(collectionInterval):
				s.start = s.end
				s.end = time.Now()
			}
		}
	}()
	return nil
}

func (s *logsReceiver) Shutdown(_ context.Context) error {
	close(s.stopperChan)
	s.wg.Wait()
	return s.client.Shutdown()
}

// parseHostNames parses out the hostname from the specified cluster host
func parseHostNames(s string, logger *zap.Logger) []string {
	if s == "" {
		return nil
	}
	elts := strings.Split(s, ",")
	hostnames := make([]string, 0, len(elts))

	for _, t := range elts {
		// separate hostname from scheme and port
		host, _, err := net.SplitHostPort(strings.TrimPrefix(t, "mongodb://"))
		if err != nil {
			logger.Error("Could not parse out hostname: " + host)
			continue
		}
		hostnames = append(hostnames, host)
	}

	return hostnames
}

// collect spins off functionality of the receiver from the Start function
func (s *logsReceiver) collect(ctx context.Context) {
	for _, projectCfg := range s.cfg.Logs.Projects {
		project, err := s.client.GetProject(ctx, projectCfg.Name)
		if err != nil {
			s.log.Error("Error retrieving project "+projectCfg.Name+":", zap.Error(err))
			continue
		}
		pc := ProjectContext{Project: *project}

		org, err := s.client.GetOrganization(ctx, project.OrgID)
		if err != nil {
			s.log.Error("Error retrieving organization", zap.Error(err))
			pc.orgName = "unknown"
		} else {
			pc.orgName = org.Name
		}

		// get clusters for each of the projects
		clusters, err := s.processClusters(ctx, *projectCfg, project.ID)
		if err != nil {
			s.log.Error("Failure to process Clusters", zap.Error(err))
		}

		s.collectClusterLogs(clusters, *projectCfg, pc)
	}
}

func (s *logsReceiver) processClusters(ctx context.Context, projectCfg LogsProjectConfig, projectID string) ([]mongodbatlas.Cluster, error) {
	clusters, err := s.client.GetClusters(ctx, projectID)
	if err != nil {
		s.log.Error("Failure to collect clusters from project: %w", zap.Error(err))
		return nil, err
	}

	return filterClusters(clusters, projectCfg.ProjectConfig)
}

func (s *logsReceiver) collectClusterLogs(clusters []mongodbatlas.Cluster, projectCfg LogsProjectConfig, pc ProjectContext) {
	for _, cluster := range clusters {
		hostnames := parseHostNames(cluster.ConnectionStrings.Standard, s.log)
		for _, hostname := range hostnames {
			// Defaults to true if not specified
			if projectCfg.EnableHostLogs == nil || *projectCfg.EnableHostLogs {
				s.collectLogs(pc, hostname, "mongodb.gz", cluster.Name, cluster.MongoDBMajorVersion)
				s.collectLogs(pc, hostname, "mongos.gz", cluster.Name, cluster.MongoDBMajorVersion)
			}

			// Defaults to false if not specified
			if projectCfg.EnableAuditLogs {
				s.collectAuditLogs(pc, hostname, "mongodb-audit-log.gz", cluster.Name, cluster.MongoDBMajorVersion)
				s.collectAuditLogs(pc, hostname, "mongos-audit-log.gz", cluster.Name, cluster.MongoDBMajorVersion)
			}
		}
	}
}

func filterClusters(clusters []mongodbatlas.Cluster, projectCfg ProjectConfig) ([]mongodbatlas.Cluster, error) {
	include, exclude := projectCfg.IncludeClusters, projectCfg.ExcludeClusters
	var whitelist bool
	var clusterNameSet map[string]struct{}
	// check to include or exclude clusters
	switch {
	// keep all clusters if include and exclude are not specified
	case len(include) == 0 && len(exclude) == 0:
		return clusters, nil
	// include is initialized
	case len(include) > 0 && len(exclude) == 0:
		whitelist = true
		clusterNameSet = projectCfg.includesByClusterName
	// exclude is initialized
	case len(exclude) > 0 && len(include) == 0:
		whitelist = false
		clusterNameSet = projectCfg.excludesByClusterName
	// both are initialized
	default:
		return nil, errors.New("both Include and Exclude clusters configured")
	}

	var filtered []mongodbatlas.Cluster
	for _, cluster := range clusters {
		if _, ok := clusterNameSet[cluster.Name]; (!ok && !whitelist) || (ok && whitelist) {
			filtered = append(filtered, cluster)
		}
	}
	return filtered, nil
}

func (s *logsReceiver) getHostLogs(groupID, hostname, logName string, clusterMajorVersion string) ([]model.LogEntry, error) {
	// Get gzip bytes buffer from API
	buf, err := s.client.GetLogs(context.Background(), groupID, hostname, logName, s.start, s.end)
	if err != nil {
		return nil, err
	}

	return decodeLogs(s.log, clusterMajorVersion, buf)
}

func (s *logsReceiver) getHostAuditLogs(groupID, hostname, logName string) ([]model.AuditLog, error) {
	// Get gzip bytes buffer from API
	buf, err := s.client.GetLogs(context.Background(), groupID, hostname, logName, s.start, s.end)
	if err != nil {
		return nil, err
	}

	return decodeAuditJSON(s.log, buf)
}

func (s *logsReceiver) collectLogs(pc ProjectContext, hostname, logName, clusterName, clusterMajorVersion string) {
	logs, err := s.getHostLogs(pc.Project.ID, hostname, logName, clusterMajorVersion)
	if err != nil && !errors.Is(err, io.EOF) {
		s.log.Warn("Failed to retrieve host logs", zap.Error(err), zap.String("log", logName))
		return
	}

	if len(logs) == 0 {
		s.log.Warn("Attempted to retrieve host logs but received 0 logs", zap.Error(err), zap.String("log", logName))
		return
	}

	plog := mongodbEventToLogData(s.log,
		logs,
		pc,
		hostname,
		logName,
		clusterName,
		clusterMajorVersion)
	err = s.consumer.ConsumeLogs(context.Background(), plog)
	if err != nil {
		s.log.Error("Failed to consume logs", zap.Error(err))
	}
}

func (s *logsReceiver) collectAuditLogs(pc ProjectContext, hostname, logName, clusterName, clusterMajorVersion string) {
	logs, err := s.getHostAuditLogs(
		pc.Project.ID,
		hostname,
		logName,
	)

	if err != nil && !errors.Is(err, io.EOF) {
		s.log.Warn("Failed to retrieve audit logs", zap.Error(err), zap.String("log", logName))
		return
	}

	if len(logs) == 0 {
		s.log.Warn("Attempted to retrieve audit logs but received 0 logs", zap.Error(err), zap.String("log", logName))
		return
	}

	plog, err := mongodbAuditEventToLogData(s.log,
		logs,
		pc,
		hostname,
		logName,
		clusterName,
		clusterMajorVersion)
	if err != nil {
		s.log.Warn("Failed to translate audit logs: "+logName, zap.Error(err))
		return
	}

	err = s.consumer.ConsumeLogs(context.Background(), plog)
	if err != nil {
		s.log.Error("Failed to consume logs", zap.Error(err))
	}
}
