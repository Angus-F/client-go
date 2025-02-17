/*
Copyright 2018 The Kubernetes Authors.

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

package dynamiclister

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/Angus-F/client-go/tools/cache"
)

var _ cache.GenericLister = &dynamicListerShim{}
var _ cache.GenericNamespaceLister = &dynamicNamespaceListerShim{}

// dynamicListerShim implements the cache.GenericLister interface.
type dynamicListerShim struct {
	lister Lister
}

// NewRuntimeObjectShim returns a new shim for Lister.
// It wraps Lister so that it implements cache.GenericLister interface
func NewRuntimeObjectShim(lister Lister) cache.GenericLister {
	return &dynamicListerShim{lister: lister}
}

// List will return all objects across namespaces
func (s *dynamicListerShim) List(selector labels.Selector) (ret []runtime.Object, err error) {
	objs, err := s.lister.List(selector)
	if err != nil {
		return nil, err
	}

	ret = make([]runtime.Object, len(objs))
	for index, obj := range objs {
		ret[index] = obj
	}
	return ret, err
}

// Get will attempt to retrieve assuming that name==key
func (s *dynamicListerShim) Get(name string) (runtime.Object, error) {
	return s.lister.Get(name)
}

func (s *dynamicListerShim) ByNamespace(namespace string) cache.GenericNamespaceLister {
	return &dynamicNamespaceListerShim{
		namespaceLister: s.lister.Namespace(namespace),
	}
}

// dynamicNamespaceListerShim implements the NamespaceLister interface.
// It wraps NamespaceLister so that it implements cache.GenericNamespaceLister interface
type dynamicNamespaceListerShim struct {
	namespaceLister NamespaceLister
}

// List will return all objects in this namespace
func (ns *dynamicNamespaceListerShim) List(selector labels.Selector) (ret []runtime.Object, err error) {
	objs, err := ns.namespaceLister.List(selector)
	if err != nil {
		return nil, err
	}

	ret = make([]runtime.Object, len(objs))
	for index, obj := range objs {
		ret[index] = obj
	}
	return ret, err
}

// Get will attempt to retrieve by namespace and name
func (ns *dynamicNamespaceListerShim) Get(name string) (runtime.Object, error) {
	return ns.namespaceLister.Get(name)
}
