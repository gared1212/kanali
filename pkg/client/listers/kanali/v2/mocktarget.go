// Copyright (c) 2017 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// This file was automatically generated by lister-gen

package v2

import (
	v2 "github.com/northwesternmutual/kanali/pkg/apis/kanali.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MockTargetLister helps list MockTargets.
type MockTargetLister interface {
	// List lists all MockTargets in the indexer.
	List(selector labels.Selector) (ret []*v2.MockTarget, err error)
	// MockTargets returns an object that can list and get MockTargets.
	MockTargets(namespace string) MockTargetNamespaceLister
	MockTargetListerExpansion
}

// mockTargetLister implements the MockTargetLister interface.
type mockTargetLister struct {
	indexer cache.Indexer
}

// NewMockTargetLister returns a new MockTargetLister.
func NewMockTargetLister(indexer cache.Indexer) MockTargetLister {
	return &mockTargetLister{indexer: indexer}
}

// List lists all MockTargets in the indexer.
func (s *mockTargetLister) List(selector labels.Selector) (ret []*v2.MockTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.MockTarget))
	})
	return ret, err
}

// MockTargets returns an object that can list and get MockTargets.
func (s *mockTargetLister) MockTargets(namespace string) MockTargetNamespaceLister {
	return mockTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MockTargetNamespaceLister helps list and get MockTargets.
type MockTargetNamespaceLister interface {
	// List lists all MockTargets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.MockTarget, err error)
	// Get retrieves the MockTarget from the indexer for a given namespace and name.
	Get(name string) (*v2.MockTarget, error)
	MockTargetNamespaceListerExpansion
}

// mockTargetNamespaceLister implements the MockTargetNamespaceLister
// interface.
type mockTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MockTargets in the indexer for a given namespace.
func (s mockTargetNamespaceLister) List(selector labels.Selector) (ret []*v2.MockTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.MockTarget))
	})
	return ret, err
}

// Get retrieves the MockTarget from the indexer for a given namespace and name.
func (s mockTargetNamespaceLister) Get(name string) (*v2.MockTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("mocktarget"), name)
	}
	return obj.(*v2.MockTarget), nil
}
