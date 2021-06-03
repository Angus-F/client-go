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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	storagev1 "k8s.io/api/storage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1 "github.com/Angus-F/client-go/applyconfigurations/storage/v1"
	testing "github.com/Angus-F/client-go/testing"
)

// FakeStorageClasses implements StorageClassInterface
type FakeStorageClasses struct {
	Fake *FakeStorageV1
}

var storageclassesResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"}

var storageclassesKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1", Kind: "StorageClass"}

// Get takes name of the storageClass, and returns the corresponding storageClass object, and an error if there is any.
func (c *FakeStorageClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *storagev1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storageclassesResource, name), &storagev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.StorageClass), err
}

// List takes label and field selectors, and returns the list of StorageClasses that match those selectors.
func (c *FakeStorageClasses) List(ctx context.Context, opts v1.ListOptions) (result *storagev1.StorageClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storageclassesResource, storageclassesKind, opts), &storagev1.StorageClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1.StorageClassList{ListMeta: obj.(*storagev1.StorageClassList).ListMeta}
	for _, item := range obj.(*storagev1.StorageClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageClasses.
func (c *FakeStorageClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storageclassesResource, opts))
}

// Create takes the representation of a storageClass and creates it.  Returns the server's representation of the storageClass, and an error, if there is any.
func (c *FakeStorageClasses) Create(ctx context.Context, storageClass *storagev1.StorageClass, opts v1.CreateOptions) (result *storagev1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storageclassesResource, storageClass), &storagev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.StorageClass), err
}

// Update takes the representation of a storageClass and updates it. Returns the server's representation of the storageClass, and an error, if there is any.
func (c *FakeStorageClasses) Update(ctx context.Context, storageClass *storagev1.StorageClass, opts v1.UpdateOptions) (result *storagev1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storageclassesResource, storageClass), &storagev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.StorageClass), err
}

// Delete takes name of the storageClass and deletes it. Returns an error if one occurs.
func (c *FakeStorageClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storageclassesResource, name), &storagev1.StorageClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storageclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1.StorageClassList{})
	return err
}

// Patch applies the patch and returns the patched storageClass.
func (c *FakeStorageClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *storagev1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storageclassesResource, name, pt, data, subresources...), &storagev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.StorageClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied storageClass.
func (c *FakeStorageClasses) Apply(ctx context.Context, storageClass *applyconfigurationsstoragev1.StorageClassApplyConfiguration, opts v1.ApplyOptions) (result *storagev1.StorageClass, err error) {
	if storageClass == nil {
		return nil, fmt.Errorf("storageClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(storageClass)
	if err != nil {
		return nil, err
	}
	name := storageClass.Name
	if name == nil {
		return nil, fmt.Errorf("storageClass.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storageclassesResource, *name, types.ApplyPatchType, data), &storagev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.StorageClass), err
}
