// Code generated by mdatagen. DO NOT EDIT.

package jmxreceiver

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("internal/poll.runtime_pollWait"))
}
