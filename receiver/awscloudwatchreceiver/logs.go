// Copyright The OpenTelemetry Authors
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

package awscloudwatchreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver"

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type logsReceiver struct {
	region              string
	profile             string
	imdsEndpoint        string
	pollInterval        time.Duration
	maxEventsPerRequest int
	lastRecordedTime    time.Time
	groupRequests       []groupRequest
	autodiscover        *AutodiscoverConfig
	logger              *zap.Logger
	client              client
	consumer            consumer.Logs
	wg                  *sync.WaitGroup
	mu                  sync.Mutex
	doneChan            chan bool
}

type client interface {
	DescribeLogGroupsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogGroupsInput, opts ...request.Option) (*cloudwatchlogs.DescribeLogGroupsOutput, error)
	FilterLogEventsWithContext(ctx context.Context, input *cloudwatchlogs.FilterLogEventsInput, opts ...request.Option) (*cloudwatchlogs.FilterLogEventsOutput, error)
}

type streamNames struct {
	group string
	names []*string
}

func (sn *streamNames) request(limit int, nextToken string, st, et *time.Time) *cloudwatchlogs.FilterLogEventsInput {
	base := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName: &sn.group,
		StartTime:    aws.Int64(st.UnixMilli()),
		EndTime:      aws.Int64(et.UnixMilli()),
		Limit:        aws.Int64(int64(limit)),
	}
	if len(sn.names) > 0 {
		base.LogStreamNames = sn.names
	}
	if nextToken != "" {
		base.NextToken = aws.String(nextToken)
	}
	return base
}

func (sn *streamNames) groupName() string {
	return sn.group
}

type streamPrefix struct {
	group  string
	prefix *string
}

func (sp *streamPrefix) request(limit int, nextToken string, st, et *time.Time) *cloudwatchlogs.FilterLogEventsInput {
	base := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName:        &sp.group,
		StartTime:           aws.Int64(st.UnixMilli()),
		EndTime:             aws.Int64(et.UnixMilli()),
		Limit:               aws.Int64(int64(limit)),
		LogStreamNamePrefix: sp.prefix,
	}
	if nextToken != "" {
		base.NextToken = aws.String(nextToken)
	}
	return base
}

func (sp *streamPrefix) groupName() string {
	return sp.group
}

type groupRequest interface {
	request(limit int, nextToken string, st, et *time.Time) *cloudwatchlogs.FilterLogEventsInput
	groupName() string
}

func newLogsReceiver(cfg *Config, logger *zap.Logger, consumer consumer.Logs) *logsReceiver {
	groups := []groupRequest{}
	for logGroupName, sc := range cfg.Logs.Groups.NamedConfigs {
		for _, prefix := range sc.Prefixes {
			groups = append(groups, &streamPrefix{group: logGroupName, prefix: prefix})
		}
		groups = append(groups, &streamNames{group: logGroupName, names: sc.Names})
	}

	// safeguard from using both
	autodiscover := cfg.Logs.Groups.AutodiscoverConfig
	if len(cfg.Logs.Groups.NamedConfigs) > 0 {
		autodiscover = nil
	}

	return &logsReceiver{
		region:              cfg.Region,
		profile:             cfg.Profile,
		consumer:            consumer,
		maxEventsPerRequest: cfg.Logs.MaxEventsPerRequest,
		imdsEndpoint:        cfg.IMDSEndpoint,
		autodiscover:        autodiscover,
		pollInterval:        cfg.Logs.PollInterval,
		lastRecordedTime:    pcommon.Timestamp(0).AsTime(),
		groupRequests:       groups,
		logger:              logger,
		wg:                  &sync.WaitGroup{},
		doneChan:            make(chan bool),
	}
}

func (l *logsReceiver) Start(ctx context.Context, host component.Host) error {
	l.logger.Debug("starting to poll for Cloudwatch logs")

	if l.autodiscover != nil {
		group, err := l.discoverGroups(ctx, l.autodiscover)
		if err != nil {
			l.logger.Error("unable to auto discover log groups using prefix", zap.Error(err))
		}
		l.groupRequests = group
		l.wg.Add(1)
		go l.startDiscovering(ctx, l.autodiscover)
	}

	l.wg.Add(1)
	go l.startPolling(ctx)
	return nil
}

func (l *logsReceiver) Shutdown(ctx context.Context) error {
	l.logger.Debug("shutting down logs receiver")
	close(l.doneChan)
	l.wg.Wait()
	return nil
}

func (l *logsReceiver) startPolling(ctx context.Context) {
	defer l.wg.Done()

	t := time.NewTicker(l.pollInterval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-l.doneChan:
			return
		case <-t.C:
			err := l.poll(ctx)
			if err != nil {
				l.logger.Error("there was an error during the poll", zap.Error(err))
			}
		}
	}
}

func (l *logsReceiver) poll(ctx context.Context) error {
	var errs error
	for _, r := range l.groupRequests {
		if err := l.pollForLogs(ctx, r); err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}

func (l *logsReceiver) pollForLogs(ctx context.Context, pc groupRequest) error {
	err := l.ensureSession()
	if err != nil {
		return err
	}
	nextToken := aws.String("")
	startTime := l.lastRecordedTime
	endTime := time.Now()
	for nextToken != nil {
		select {
		// if done, we want to stop processing paginated stream of events
		case _, ok := <-l.doneChan:
			if !ok {
				return nil
			}
		default:
			input := pc.request(l.maxEventsPerRequest, *nextToken, &startTime, &endTime)
			resp, err := l.client.FilterLogEventsWithContext(ctx, input)
			if err != nil {
				l.logger.Error("unable to retrieve logs from cloudwatch", zap.String("log group", pc.groupName()), zap.Error(err))
				break
			}
			observedTime := pcommon.NewTimestampFromTime(time.Now())
			logs := l.processEvents(observedTime, pc.groupName(), resp)
			if logs.LogRecordCount() > 0 {
				if err = l.consumer.ConsumeLogs(ctx, logs); err != nil {
					l.logger.Error("unable to consume logs", zap.Error(err))
					break
				}
			}
			nextToken = resp.NextToken
		}
	}
	return nil
}

func (l *logsReceiver) processEvents(now pcommon.Timestamp, logGroupName string, output *cloudwatchlogs.FilterLogEventsOutput) plog.Logs {
	logs := plog.NewLogs()
	for _, e := range output.Events {
		if e.Timestamp == nil {
			l.logger.Error("unable to determine timestamp of event as the timestamp is nil")
			continue
		}

		if e.EventId == nil {
			l.logger.Error("no event ID was present on the event, skipping entry")
			continue
		}

		if e.Message == nil {
			l.logger.Error("no message was present on the event", zap.String("event.id", *e.EventId))
			continue
		}

		rl := logs.ResourceLogs().AppendEmpty()
		resourceAttributes := rl.Resource().Attributes()
		resourceAttributes.PutString("aws.region", l.region)
		resourceAttributes.PutString("cloudwatch.log.group.name", logGroupName)
		if e.LogStreamName != nil {
			resourceAttributes.PutString("cloudwatch.log.stream", *e.LogStreamName)
		}

		logRecord := rl.ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()
		logRecord.SetObservedTimestamp(now)
		ts := time.UnixMilli(*e.Timestamp)
		logRecord.SetTimestamp(pcommon.NewTimestampFromTime(ts))

		if ts.After(l.lastRecordedTime) {
			l.lastRecordedTime = ts
		}

		logRecord.Body().SetStr(*e.Message)
		logRecord.Attributes().PutString("id", *e.EventId)
	}
	return logs
}

func (l *logsReceiver) startDiscovering(ctx context.Context, auto *AutodiscoverConfig) {
	defer l.wg.Done()
	t := time.NewTicker(l.pollInterval)
	for {
		select {
		case _, ok := <-l.doneChan:
			if !ok {
				return
			}
		case <-ctx.Done():
			return
		case <-t.C:
			group, err := l.discoverGroups(ctx, auto)
			if err != nil {
				l.logger.Error("unable to perform discovery of log groups", zap.Error(err))
			}
			l.groupRequests = group
		}
	}
}

func (l *logsReceiver) discoverGroups(ctx context.Context, auto *AutodiscoverConfig) ([]groupRequest, error) {
	l.logger.Debug("attempting to discover log groups.", zap.Int("limit", auto.Limit))
	groups := []groupRequest{}
	err := l.ensureSession()
	if err != nil {
		return groups, fmt.Errorf("unable to establish a session to auto discover log groups: %w", err)
	}

	numGroups := 0
	var nextToken = aws.String("")
	for nextToken != nil {
		if numGroups > auto.Limit {
			break
		}

		limit := int64(auto.Limit)
		req := &cloudwatchlogs.DescribeLogGroupsInput{
			Limit: aws.Int64(limit),
		}

		if auto.Prefix != "" {
			req.LogGroupNamePrefix = &auto.Prefix
		}

		dlgResults, err := l.client.DescribeLogGroupsWithContext(ctx, req)
		if err != nil {
			return groups, fmt.Errorf("unable to list log groups: %w", err)
		}

		for _, lg := range dlgResults.LogGroups {
			numGroups++
			l.logger.Debug("discovered log group", zap.String("log group", lg.GoString()))
			// default behavior is to collect all if not stream filtered
			if len(auto.Streams.Names) == 0 && len(auto.Streams.Prefixes) == 0 {
				groups = append(groups, &streamNames{group: *lg.LogGroupName})
				continue
			}

			for _, prefix := range auto.Streams.Prefixes {
				groups = append(groups, &streamPrefix{group: *lg.LogGroupName, prefix: prefix})
			}

			if len(auto.Streams.Names) > 0 {
				groups = append(groups, &streamNames{group: *lg.LogGroupName, names: auto.Streams.Names})
			}
		}
		nextToken = dlgResults.NextToken
	}

	return groups, nil
}

func (l *logsReceiver) ensureSession() error {
	if l.client != nil {
		return nil
	}
	awsConfig := aws.NewConfig().WithRegion(l.region)
	options := session.Options{
		Config: *awsConfig,
	}
	if l.imdsEndpoint != "" {
		options.EC2IMDSEndpoint = l.imdsEndpoint
	}
	if l.profile != "" {
		options.Profile = l.profile
	}
	s, err := session.NewSessionWithOptions(options)
	l.client = cloudwatchlogs.New(s)
	return err
}
