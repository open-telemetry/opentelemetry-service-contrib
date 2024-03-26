// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8stest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8stest"

import (
	"errors"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/util/net"
	"k8s.io/client-go/discovery"
	memory "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	DynamicClient   *dynamic.DynamicClient
	DiscoveryClient *discovery.DiscoveryClient
	Mapper          *restmapper.DeferredDiscoveryRESTMapper

	httpClient *http.Client
}

func NewK8sClient(kubeconfigPath string) (*K8sClient, error) {
	if kubeconfigPath == "" {
		return nil, errors.New("Please provide file path to load kubeconfig")
	}
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load kubeconfig from %s: %w", kubeconfigPath, err)
	}

	restConfig.Proxy = http.ProxyFromEnvironment
	httpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return nil, err
	}

	dynamicClient, err := dynamic.NewForConfigAndClient(restConfig, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error creating dynamic client: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating discovery client: %w", err)
	}

	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))

	return &K8sClient{
		DynamicClient: dynamicClient, DiscoveryClient: discoveryClient, Mapper: mapper, httpClient: httpClient}, nil
}

func (k *K8sClient) Shutdown() {
	if k.httpClient != nil {
		net.CloseIdleConnectionsFor(k.httpClient.Transport)
	}
	return
}
