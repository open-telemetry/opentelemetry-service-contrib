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

// +build linux

package pagingscraper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const swapsFilePath = "/proc/swaps"

// swaps file column indexes
const (
	nameCol = 0
	// typeCol     = 1
	totalCol = 2
	usedCol  = 3
	// priorityCol = 4
)

func getPageFileStats() ([]*pageFileStats, error) {
	f, err := os.Open(swapsFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parseSwapsFile(f)
}

func parseSwapsFile(r io.Reader) ([]*pageFileStats, error) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("couldn't read file %q: %w", swapsFilePath, err)
		}
		return nil, fmt.Errorf("unexpected end-of-file in %q", swapsFilePath)

	}

	// Check header headerFields are as expected
	headerFields := strings.Fields(scanner.Text())
	if len(headerFields) < usedCol {
		return nil, fmt.Errorf("couldn't parse %q: too few fields in header", swapsFilePath)
	}
	if headerFields[nameCol] != "Filename" {
		return nil, fmt.Errorf("couldn't parse %q: expected %q to be %q", swapsFilePath, headerFields[nameCol], "Filename")
	}
	if headerFields[totalCol] != "Size" {
		return nil, fmt.Errorf("couldn't parse %q: expected %q to be %q", swapsFilePath, headerFields[totalCol], "Size")
	}
	if headerFields[usedCol] != "Used" {
		return nil, fmt.Errorf("couldn't parse %q: expected %q to be %q", swapsFilePath, headerFields[usedCol], "Used")
	}

	var swapDevices []*pageFileStats
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < usedCol {
			return nil, fmt.Errorf("couldn't parse %q: too few fields", swapsFilePath)
		}

		totalKiB, err := strconv.ParseUint(fields[totalCol], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse 'Size' column in %q: %w", swapsFilePath, err)
		}

		usedKiB, err := strconv.ParseUint(fields[usedCol], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse 'Used' column in %q: %w", swapsFilePath, err)
		}

		swapDevices = append(swapDevices, &pageFileStats{
			deviceName: fields[nameCol],
			usedBytes:  usedKiB * 1024,
			freeBytes:  (totalKiB - usedKiB) * 1024,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("couldn't read file %q: %w", swapsFilePath, err)
	}

	return swapDevices, nil
}
