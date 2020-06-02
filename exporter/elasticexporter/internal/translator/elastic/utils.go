// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package elastic contains an opentelemetry-collector exporter
// for Elastic APM.
package elastic

import (
	"fmt"
	"regexp"
	"strings"

	"go.opentelemetry.io/collector/consumer/pdata"
)

var (
	serviceNameInvalidRegexp = regexp.MustCompile("[^a-zA-Z0-9 _-]")
	labelKeyReplacer         = strings.NewReplacer(`.`, `_`, `*`, `_`, `"`, `_`)
)

func combineErrors(errs []error) error {
	if n := len(errs); n == 0 {
		return nil
	} else if n == 1 {
		return errs[0]
	}
	return fmt.Errorf("%q", errs)
}

func ifaceAttributeValue(v pdata.AttributeValue) interface{} {
	switch v.Type() {
	case pdata.AttributeValueSTRING:
		return truncate(v.StringVal())
	case pdata.AttributeValueINT:
		return v.IntVal()
	case pdata.AttributeValueDOUBLE:
		return v.DoubleVal()
	case pdata.AttributeValueBOOL:
		return v.BoolVal()
	}
	return nil
}

func cleanServiceName(name string) string {
	return serviceNameInvalidRegexp.ReplaceAllString(truncate(name), "_")
}

func cleanLabelKey(k string) string {
	return labelKeyReplacer.Replace(truncate(k))
}

func truncate(s string) string {
	const maxRunes = 1024
	var j int
	for i := range s {
		if j == maxRunes {
			return s[:i]
		}
		j++
	}
	return s
}
