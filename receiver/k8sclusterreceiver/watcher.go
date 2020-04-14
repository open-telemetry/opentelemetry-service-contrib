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

package k8sclusterreceiver

import (
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/autoscaling/v2beta1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type resourceWatcher struct {
	client                kubernetes.Interface
	sharedInformerFactory informers.SharedInformerFactory
	dataCollector         *dataCollector
	logger                *zap.Logger
	// This field is temporary and will be removed once the
	// metadata syncing details are finalized.
	collectMedata bool
}

// newResourceWatcher creates a Kubernetes resource watcher.
func newResourceWatcher(logger *zap.Logger, config *Config,
	client kubernetes.Interface, collectMetadata bool) (*resourceWatcher, error) {
	rw := &resourceWatcher{
		client:        client,
		logger:        logger,
		dataCollector: newDataCollector(logger, config.NodeConditionTypesToReport),
		collectMedata: collectMetadata,
	}

	rw.prepareSharedInformerFactory()

	return rw, nil
}

func (rw *resourceWatcher) prepareSharedInformerFactory() {
	factory := informers.NewSharedInformerFactoryWithOptions(rw.client, 0)

	// Add shared informers for each resource type that has to be watched.
	rw.setupInformers(&corev1.Pod{}, factory.Core().V1().Pods().Informer())
	rw.setupInformers(&corev1.Node{}, factory.Core().V1().Nodes().Informer())
	rw.setupInformers(&corev1.Namespace{}, factory.Core().V1().Namespaces().Informer())
	rw.setupInformers(&corev1.ReplicationController{},
		factory.Core().V1().ReplicationControllers().Informer(),
	)
	rw.setupInformers(&corev1.ResourceQuota{}, factory.Core().V1().ResourceQuotas().Informer())
	// rw.setupInformers(&corev1.Service{}, factory.Core().V1().Services().Informer())
	rw.setupInformers(&appsv1.DaemonSet{}, factory.Apps().V1().DaemonSets().Informer())
	rw.setupInformers(&appsv1.Deployment{}, factory.Apps().V1().Deployments().Informer())
	rw.setupInformers(&appsv1.ReplicaSet{}, factory.Apps().V1().ReplicaSets().Informer())
	rw.setupInformers(&appsv1.StatefulSet{}, factory.Apps().V1().StatefulSets().Informer())
	rw.setupInformers(&batchv1.Job{}, factory.Batch().V1().Jobs().Informer())
	rw.setupInformers(&batchv1beta1.CronJob{}, factory.Batch().V1beta1().CronJobs().Informer())
	rw.setupInformers(&v2beta1.HorizontalPodAutoscaler{},
		factory.Autoscaling().V2beta1().HorizontalPodAutoscalers().Informer(),
	)

	rw.sharedInformerFactory = factory
}

// startWatchingResources starts up all informers.
func (rw *resourceWatcher) startWatchingResources(stopper <-chan struct{}) {
	rw.sharedInformerFactory.Start(stopper)
}

// setupInformers adds event handlers to informers and setups a metadataStore.
func (rw *resourceWatcher) setupInformers(o runtime.Object, informer cache.SharedIndexInformer) {
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    rw.onAdd,
		UpdateFunc: rw.onUpdate,
		DeleteFunc: rw.onDelete,
	})
	rw.dataCollector.setupMetadataStore(o, informer.GetStore())
}

func (rw *resourceWatcher) onAdd(obj interface{}) {
	rw.addOrUpdateResource(obj)
}

func (rw *resourceWatcher) onDelete(obj interface{}) {
	rw.dataCollector.removeFromMetricsStore(obj)
}

func (rw *resourceWatcher) onUpdate(_ interface{}, newObj interface{}) {
	rw.addOrUpdateResource(newObj)
}

// addOrUpdateResource keeps the metric cache updated with latest metric
// values that will be dispatched in the subsequent interval.
func (rw *resourceWatcher) addOrUpdateResource(obj interface{}) {
	rw.dataCollector.syncMetrics(obj)

	// Sync metadata only if there's at least one destination for it to sent.
	if rw.collectMedata {
		rw.dataCollector.syncMetadata(obj)
	}
}
