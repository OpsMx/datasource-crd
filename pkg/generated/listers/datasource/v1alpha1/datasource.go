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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/opsmx/datasources-crd/pkg/apis/datasource/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DatasourceLister helps list Datasources.
// All objects returned here must be treated as read-only.
type DatasourceLister interface {
	// List lists all Datasources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Datasource, err error)
	// Datasources returns an object that can list and get Datasources.
	Datasources(namespace string) DatasourceNamespaceLister
	DatasourceListerExpansion
}

// datasourceLister implements the DatasourceLister interface.
type datasourceLister struct {
	indexer cache.Indexer
}

// NewDatasourceLister returns a new DatasourceLister.
func NewDatasourceLister(indexer cache.Indexer) DatasourceLister {
	return &datasourceLister{indexer: indexer}
}

// List lists all Datasources in the indexer.
func (s *datasourceLister) List(selector labels.Selector) (ret []*v1alpha1.Datasource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Datasource))
	})
	return ret, err
}

// Datasources returns an object that can list and get Datasources.
func (s *datasourceLister) Datasources(namespace string) DatasourceNamespaceLister {
	return datasourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DatasourceNamespaceLister helps list and get Datasources.
// All objects returned here must be treated as read-only.
type DatasourceNamespaceLister interface {
	// List lists all Datasources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Datasource, err error)
	// Get retrieves the Datasource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Datasource, error)
	DatasourceNamespaceListerExpansion
}

// datasourceNamespaceLister implements the DatasourceNamespaceLister
// interface.
type datasourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Datasources in the indexer for a given namespace.
func (s datasourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Datasource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Datasource))
	})
	return ret, err
}

// Get retrieves the Datasource from the indexer for a given namespace and name.
func (s datasourceNamespaceLister) Get(name string) (*v1alpha1.Datasource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datasource"), name)
	}
	return obj.(*v1alpha1.Datasource), nil
}
