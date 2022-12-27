/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	datasourcev1alpha1 "github.com/opsmx/datasource-crd/pkg/apis/datasource/v1alpha1"
	versioned "github.com/opsmx/datasource-crd/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/opsmx/datasource-crd/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/opsmx/datasource-crd/pkg/generated/listers/datasource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DatasourceInformer provides access to a shared informer and lister for
// Datasources.
type DatasourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DatasourceLister
}

type datasourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDatasourceInformer constructs a new informer for Datasource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDatasourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDatasourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDatasourceInformer constructs a new informer for Datasource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDatasourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AnalyticsV1alpha1().Datasources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AnalyticsV1alpha1().Datasources(namespace).Watch(context.TODO(), options)
			},
		},
		&datasourcev1alpha1.Datasource{},
		resyncPeriod,
		indexers,
	)
}

func (f *datasourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDatasourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *datasourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&datasourcev1alpha1.Datasource{}, f.defaultInformer)
}

func (f *datasourceInformer) Lister() v1alpha1.DatasourceLister {
	return v1alpha1.NewDatasourceLister(f.Informer().GetIndexer())
}
