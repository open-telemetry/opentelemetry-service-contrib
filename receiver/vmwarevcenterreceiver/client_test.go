package vmwarevcenterreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func TestGetClusters(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)
		client := vmwareVcenterClient{
			vimDriver: c,
			finder:    finder,
		}
		clusters, err := client.Clusters(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, clusters, 0)
	})
}

func TestGetResourcePools(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)
		client := vmwareVcenterClient{
			vimDriver: c,
			finder:    finder,
		}
		resourcePools, err := client.ResourcePools(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, resourcePools)
	})
}

func TestGetVMs(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)
		client := vmwareVcenterClient{
			vimDriver: c,
			finder:    finder,
		}
		vms, err := client.VMs(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, vms)
	})
}
