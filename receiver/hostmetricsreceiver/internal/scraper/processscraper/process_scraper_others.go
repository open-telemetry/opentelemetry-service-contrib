// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !linux && !windows
// +build !linux,!windows

package processscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper/internal/metadata"
)

func (s *scraper) recordCPUTimeMetric(now pcommon.Timestamp, cpuTime *cpu.TimesStat) {
	s.mb.RecordProcessCPUTimeDataPoint(now, cpuTime.User, metadata.AttributeStateUser)
	s.mb.RecordProcessCPUTimeDataPoint(now, cpuTime.System, metadata.AttributeStateSystem)
	s.mb.RecordProcessCPUTimeDataPoint(now, cpuTime.Iowait, metadata.AttributeStateWait)
}

func getProcessExecutable(processHandle) (*executableMetadata, error) {
	// attempting to get the process name 
	name, err := proc.Name()
	if err != nil {
		return nil, err
	}

	// attampting to get the process path
	exe, err := proc.Exe()
	if err != nil {
		return nil, err
	}

	// managed to get the info
	executable := &executableMetadata{name: name, path: exe}
	return executable, nil
}

func getProcessCommand(processHandle) (*commandMetadata, error) {
	cmdline, err := proc.Cmdline()
	if err != nil {
		return nil, err
	}

	command := &commandMetadata{command: cmdline, commandLine: cmdline}
	return command, nil

}
