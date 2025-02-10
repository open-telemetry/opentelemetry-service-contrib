// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package webhookeventreceiver

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
	"net/textproto"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receivertest"
)

func TestReqToLog(t *testing.T) {
	defaultConfig := createDefaultConfig().(*Config)

	tests := []struct {
		desc    string
		sc      *bufio.Scanner
		headers http.Header
		query   url.Values
		config  *Config
		tt      func(t *testing.T, reqLog plog.Logs, reqLen int, settings receiver.Settings)
	}{
		{
			desc: "Valid query valid event",
			sc: func() *bufio.Scanner {
				reader := io.NopCloser(bytes.NewReader([]byte("this is a: log")))
				return bufio.NewScanner(reader)
			}(),
			query: func() url.Values {
				v, err := url.ParseQuery(`qparam1=hello&qparam2=world`)
				if err != nil {
					log.Fatal("failed to parse query")
				}
				return v
			}(),
			tt: func(t *testing.T, reqLog plog.Logs, reqLen int, _ receiver.Settings) {
				require.Equal(t, 1, reqLen)

				attributes := reqLog.ResourceLogs().At(0).Resource().Attributes()
				require.Equal(t, 2, attributes.Len())

				scopeLogsScope := reqLog.ResourceLogs().At(0).ScopeLogs().At(0).Scope()
				require.Equal(t, 2, scopeLogsScope.Attributes().Len())

				if v, ok := attributes.Get("qparam1"); ok {
					require.Equal(t, "hello", v.AsString())
				} else {
					require.Fail(t, "failed to set attribute from query parameter 1")
				}
				if v, ok := attributes.Get("qparam2"); ok {
					require.Equal(t, "world", v.AsString())
				} else {
					require.Fail(t, "failed to set attribute query parameter 2")
				}
			},
		},
		{
			desc: "new lines present in body",
			sc: func() *bufio.Scanner {
				reader := io.NopCloser(bytes.NewReader([]byte("{\n\"key\":\"value\"\n}")))
				return bufio.NewScanner(reader)
			}(),
			query: func() url.Values {
				v, err := url.ParseQuery(`qparam1=hello&qparam2=world`)
				if err != nil {
					log.Fatal("failed to parse query")
				}
				return v
			}(),
			tt: func(t *testing.T, reqLog plog.Logs, reqLen int, _ receiver.Settings) {
				require.Equal(t, 1, reqLen)

				attributes := reqLog.ResourceLogs().At(0).Resource().Attributes()
				require.Equal(t, 2, attributes.Len())

				scopeLogsScope := reqLog.ResourceLogs().At(0).ScopeLogs().At(0).Scope()
				require.Equal(t, 2, scopeLogsScope.Attributes().Len())

				if v, ok := attributes.Get("qparam1"); ok {
					require.Equal(t, "hello", v.AsString())
				} else {
					require.Fail(t, "failed to set attribute from query parameter 1")
				}
				if v, ok := attributes.Get("qparam2"); ok {
					require.Equal(t, "world", v.AsString())
				} else {
					require.Fail(t, "failed to set attribute query parameter 2")
				}
			},
		},
		{
			desc: "Query is empty",
			sc: func() *bufio.Scanner {
				reader := io.NopCloser(bytes.NewReader([]byte("this is a: log")))
				return bufio.NewScanner(reader)
			}(),
			tt: func(t *testing.T, reqLog plog.Logs, reqLen int, _ receiver.Settings) {
				require.Equal(t, 1, reqLen)

				attributes := reqLog.ResourceLogs().At(0).Resource().Attributes()
				require.Equal(t, 0, attributes.Len())

				scopeLogsScope := reqLog.ResourceLogs().At(0).ScopeLogs().At(0).Scope()
				require.Equal(t, 2, scopeLogsScope.Attributes().Len())
			},
		},
		{
			desc: "Headers not added by default",
			headers: http.Header{
				textproto.CanonicalMIMEHeaderKey("X-Foo"): []string{"1"},
				textproto.CanonicalMIMEHeaderKey("X-Bar"): []string{"2"},
			},
			sc: func() *bufio.Scanner {
				reader := io.NopCloser(bytes.NewReader([]byte("this is a: log")))
				return bufio.NewScanner(reader)
			}(),
			tt: func(t *testing.T, reqLog plog.Logs, reqLen int, _ receiver.Settings) {
				require.Equal(t, 1, reqLen)

				attributes := reqLog.ResourceLogs().At(0).Resource().Attributes()
				require.Equal(t, 0, attributes.Len())

				scopeLogsScope := reqLog.ResourceLogs().At(0).ScopeLogs().At(0).Scope()
				require.Equal(t, 2, scopeLogsScope.Attributes().Len()) // expect no additional attributes even though headers are set
			},
		},
		{
			desc: "Headers added if ConvertHeadersToAttributes enabled",
			headers: http.Header{
				textproto.CanonicalMIMEHeaderKey("X-Foo"): []string{"1"},
				textproto.CanonicalMIMEHeaderKey("X-Bar"): []string{"2"},
			},
			config: &Config{
				Path:                       defaultPath,
				HealthPath:                 defaultHealthPath,
				ReadTimeout:                defaultReadTimeout,
				WriteTimeout:               defaultWriteTimeout,
				ConvertHeadersToAttributes: true,
			},
			sc: func() *bufio.Scanner {
				reader := io.NopCloser(bytes.NewReader([]byte("this is a: log")))
				return bufio.NewScanner(reader)
			}(),
			tt: func(t *testing.T, reqLog plog.Logs, reqLen int, _ receiver.Settings) {
				require.Equal(t, 1, reqLen)

				attributes := reqLog.ResourceLogs().At(0).Resource().Attributes()
				require.Equal(t, 0, attributes.Len())

				scopeLogsScope := reqLog.ResourceLogs().At(0).ScopeLogs().At(0).Scope()
				require.Equal(t, 4, scopeLogsScope.Attributes().Len()) // expect no additional attributes even though headers are set
				v, exists := scopeLogsScope.Attributes().Get("header.x_foo")
				require.True(t, exists)
				require.Equal(t, "1", v.AsString())
				v, exists = scopeLogsScope.Attributes().Get("header.x_bar")
				require.True(t, exists)
				require.Equal(t, "2", v.AsString())
			},
		},
		{
			desc: "Required header skipped",
			headers: http.Header{
				textproto.CanonicalMIMEHeaderKey("X-Foo"):             []string{"1"},
				textproto.CanonicalMIMEHeaderKey("X-Bar"):             []string{"2"},
				textproto.CanonicalMIMEHeaderKey("X-Required-Header"): []string{"password"},
			},
			config: &Config{
				Path:                       defaultPath,
				HealthPath:                 defaultHealthPath,
				ReadTimeout:                defaultReadTimeout,
				WriteTimeout:               defaultWriteTimeout,
				RequiredHeader:             RequiredHeader{Key: "X-Required-Header", Value: "password"},
				ConvertHeadersToAttributes: true,
			},
			sc: func() *bufio.Scanner {
				reader := io.NopCloser(bytes.NewReader([]byte("this is a: log")))
				return bufio.NewScanner(reader)
			}(),
			tt: func(t *testing.T, reqLog plog.Logs, reqLen int, _ receiver.Settings) {
				require.Equal(t, 1, reqLen)

				attributes := reqLog.ResourceLogs().At(0).Resource().Attributes()
				require.Equal(t, 0, attributes.Len())

				scopeLogsScope := reqLog.ResourceLogs().At(0).ScopeLogs().At(0).Scope()
				require.Equal(t, 4, scopeLogsScope.Attributes().Len()) // expect no additional attributes even though headers are set
				_, exists := scopeLogsScope.Attributes().Get("header.x_foo")
				require.True(t, exists)
				_, exists = scopeLogsScope.Attributes().Get("header.x_bar")
				require.True(t, exists)
				_, exists = scopeLogsScope.Attributes().Get("header.x_required_header")
				require.False(t, exists)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			testConfig := defaultConfig
			if test.config != nil {
				testConfig = test.config
			}
			reqLog, reqLen := reqToLog(test.sc, test.headers, test.query, testConfig, receivertest.NewNopSettings())
			test.tt(t, reqLog, reqLen, receivertest.NewNopSettings())
		})
	}
}
