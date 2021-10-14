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

package internalversion

import (
	v1alpha1 "github.com/Revolyssup/kube-client-go/pkg/apis/goglot.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GlotPodLister helps list GlotPods.
// All objects returned here must be treated as read-only.
type GlotPodLister interface {
	// List lists all GlotPods in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GlotPod, err error)
	// GlotPods returns an object that can list and get GlotPods.
	GlotPods(namespace string) GlotPodNamespaceLister
	GlotPodListerExpansion
}

// glotPodLister implements the GlotPodLister interface.
type glotPodLister struct {
	indexer cache.Indexer
}

// NewGlotPodLister returns a new GlotPodLister.
func NewGlotPodLister(indexer cache.Indexer) GlotPodLister {
	return &glotPodLister{indexer: indexer}
}

// List lists all GlotPods in the indexer.
func (s *glotPodLister) List(selector labels.Selector) (ret []*v1alpha1.GlotPod, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlotPod))
	})
	return ret, err
}

// GlotPods returns an object that can list and get GlotPods.
func (s *glotPodLister) GlotPods(namespace string) GlotPodNamespaceLister {
	return glotPodNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlotPodNamespaceLister helps list and get GlotPods.
// All objects returned here must be treated as read-only.
type GlotPodNamespaceLister interface {
	// List lists all GlotPods in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GlotPod, err error)
	// Get retrieves the GlotPod from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GlotPod, error)
	GlotPodNamespaceListerExpansion
}

// glotPodNamespaceLister implements the GlotPodNamespaceLister
// interface.
type glotPodNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlotPods in the indexer for a given namespace.
func (s glotPodNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GlotPod, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlotPod))
	})
	return ret, err
}

// Get retrieves the GlotPod from the indexer for a given namespace and name.
func (s glotPodNamespaceLister) Get(name string) (*v1alpha1.GlotPod, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("glotpod"), name)
	}
	return obj.(*v1alpha1.GlotPod), nil
}
