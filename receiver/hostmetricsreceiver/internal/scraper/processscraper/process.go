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

package processscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"

import (
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper/internal/metadata"
)

// processMetadata stores process related metadata along
// with the process handle, and provides a function to
// initialize a pdata.Resource with the metadata

type processMetadata struct {
	pid        int32
	executable *executableMetadata
	command    *commandMetadata
	username   string
	handle     processHandle
}

type executableMetadata struct {
	name string
	path string
}

type commandMetadata struct {
	command          string
	commandLine      string
	commandLineSlice []string
}

func (m *processMetadata) resourceOptions() []metadata.ResourceOption {
	opts := make([]metadata.ResourceOption, 0, 6)
	opts = append(opts,
		metadata.WithProcessPid(int64(m.pid)),
		metadata.WithProcessExecutableName(m.executable.name),
		metadata.WithProcessExecutablePath(m.executable.path),
	)
	if m.command != nil {
		opts = append(opts, metadata.WithProcessCommand(m.command.command))
		if m.command.commandLineSlice != nil {
			// TODO insert slice here once this is supported by the data model
			// (see https://github.com/open-telemetry/opentelemetry-collector/pull/1142)
			opts = append(opts, metadata.WithProcessCommandLine(strings.Join(m.command.commandLineSlice, " ")))
		} else {
			opts = append(opts, metadata.WithProcessCommandLine(m.command.commandLine))
		}
	}
	if m.username != "" {
		opts = append(opts, metadata.WithProcessOwner(m.username))
	}
	return opts
}

// processHandles provides a wrapper around []*process.Process
// to support testing

type processHandles interface {
	Pid(index int) int32
	At(index int) processHandle
	Len() int
}

type processHandle interface {
	Name() (string, error)
	Exe() (string, error)
	Username() (string, error)
	Cmdline() (string, error)
	CmdlineSlice() ([]string, error)
	Times() (*cpu.TimesStat, error)
	MemoryInfo() (*process.MemoryInfoStat, error)
	IOCounters() (*process.IOCountersStat, error)
}

type gopsProcessHandles struct {
	handles []*process.Process
}

func (p *gopsProcessHandles) Pid(index int) int32 {
	return p.handles[index].Pid
}

func (p *gopsProcessHandles) At(index int) processHandle {
	return p.handles[index]
}

func (p *gopsProcessHandles) Len() int {
	return len(p.handles)
}

func getProcessHandlesInternal() (processHandles, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	return &gopsProcessHandles{handles: processes}, nil
}
