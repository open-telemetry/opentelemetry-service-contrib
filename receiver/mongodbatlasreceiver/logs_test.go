package mongodbatlasreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
)

func TestFilterHostName(t *testing.T) {
	tmp := "mongodb://cluster0-shard-00-00.t5hdg.mongodb.net:27017,cluster0-shard-00-01.t5hdg.mongodb.net:27017,cluster0-shard-00-02.t5hdg.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=atlas-zx8u63-shard-0"
	hostnames := FilterHostName(tmp)
	require.Equal(t, []string{"cluster0-shard-00-00.t5hdg.mongodb.net", "cluster0-shard-00-01.t5hdg.mongodb.net", "cluster0-shard-00-02.t5hdg.mongodb.net"}, hostnames)
}

func TestFilterClusters(t *testing.T) {
	clusters := []mongodbatlas.Cluster{{Name: "cluster1", ID: "1"}, {Name: "cluster2", ID: "2"}, {Name: "cluster3", ID: "3"}}

	exclude := []string{"cluster1", "cluster3"}
	include := []string{"cluster1", "cluster3"}
	ec, err := filterClusters(clusters, createStringSet(exclude), false)
	require.NoError(t, err)
	require.Equal(t, []mongodbatlas.Cluster{{Name: "cluster2", ID: "2"}}, ec)

	ic, err := filterClusters(clusters, createStringSet(include), true)
	require.NoError(t, err)
	require.Equal(t, []mongodbatlas.Cluster{{Name: "cluster1", ID: "1"}, {Name: "cluster3", ID: "3"}}, ic)

}

func TestDefaultLoggingConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	params := componenttest.NewNopReceiverCreateSettings()
	ctx := context.Background()

	receiver, err := createCombinedLogReceiver(
		ctx,
		params,
		cfg,
		consumertest.NewNop(),
	)
	require.NoError(t, err)
	require.NotNil(t, receiver, "receiver creation failed")

	err = receiver.Start(ctx, componenttest.NewNopHost())
	require.NoError(t, err)

	err = receiver.Shutdown(ctx)
	require.NoError(t, err)
}
