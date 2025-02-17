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

package v1

import (
	"context"
	time "time"

	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/Angus-F/client-go/informers/internalinterfaces"
	kubernetes "github.com/Angus-F/client-go/kubernetes"
	v1 "github.com/Angus-F/client-go/listers/storage/v1"
	cache "github.com/Angus-F/client-go/tools/cache"
)

// CSINodeInformer provides access to a shared informer and lister for
// CSINodes.
type CSINodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CSINodeLister
}

type cSINodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCSINodeInformer constructs a new informer for CSINode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCSINodeInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCSINodeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCSINodeInformer constructs a new informer for CSINode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCSINodeInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().CSINodes().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().CSINodes().Watch(context.TODO(), options)
			},
		},
		&storagev1.CSINode{},
		resyncPeriod,
		indexers,
	)
}

func (f *cSINodeInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCSINodeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cSINodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1.CSINode{}, f.defaultInformer)
}

func (f *cSINodeInformer) Lister() v1.CSINodeLister {
	return v1.NewCSINodeLister(f.Informer().GetIndexer())
}
