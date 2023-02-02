/*
Copyright 2023 The Kubernetes Authors.

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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeResourceTopologyLister helps list NodeResourceTopologies.
// All objects returned here must be treated as read-only.
type NodeResourceTopologyLister interface {
	// List lists all NodeResourceTopologies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeResourceTopology, err error)
	// Get retrieves the NodeResourceTopology from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NodeResourceTopology, error)
	NodeResourceTopologyListerExpansion
}

// nodeResourceTopologyLister implements the NodeResourceTopologyLister interface.
type nodeResourceTopologyLister struct {
	indexer cache.Indexer
}

// NewNodeResourceTopologyLister returns a new NodeResourceTopologyLister.
func NewNodeResourceTopologyLister(indexer cache.Indexer) NodeResourceTopologyLister {
	return &nodeResourceTopologyLister{indexer: indexer}
}

// List lists all NodeResourceTopologies in the indexer.
func (s *nodeResourceTopologyLister) List(selector labels.Selector) (ret []*v1alpha1.NodeResourceTopology, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeResourceTopology))
	})
	return ret, err
}

// Get retrieves the NodeResourceTopology from the index for a given name.
func (s *nodeResourceTopologyLister) Get(name string) (*v1alpha1.NodeResourceTopology, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("noderesourcetopology"), name)
	}
	return obj.(*v1alpha1.NodeResourceTopology), nil
}
