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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apiautoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "github.com/Angus-F/client-go/applyconfigurations/internal"
	v1 "github.com/Angus-F/client-go/applyconfigurations/meta/v1"
)

// HorizontalPodAutoscalerApplyConfiguration represents an declarative configuration of the HorizontalPodAutoscaler type for use
// with apply.
type HorizontalPodAutoscalerApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *HorizontalPodAutoscalerSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *HorizontalPodAutoscalerStatusApplyConfiguration `json:"status,omitempty"`
}

// HorizontalPodAutoscaler constructs an declarative configuration of the HorizontalPodAutoscaler type for use with
// apply.
func HorizontalPodAutoscaler(name, namespace string) *HorizontalPodAutoscalerApplyConfiguration {
	b := &HorizontalPodAutoscalerApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("HorizontalPodAutoscaler")
	b.WithAPIVersion("autoscaling/v1")
	return b
}

// ExtractHorizontalPodAutoscaler extracts the applied configuration owned by fieldManager from
// horizontalPodAutoscaler. If no managedFields are found in horizontalPodAutoscaler for fieldManager, a
// HorizontalPodAutoscalerApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. Is is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// horizontalPodAutoscaler must be a unmodified HorizontalPodAutoscaler API object that was retrieved from the Kubernetes API.
// ExtractHorizontalPodAutoscaler provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractHorizontalPodAutoscaler(horizontalPodAutoscaler *apiautoscalingv1.HorizontalPodAutoscaler, fieldManager string) (*HorizontalPodAutoscalerApplyConfiguration, error) {
	return extractHorizontalPodAutoscaler(horizontalPodAutoscaler, fieldManager, "")
}

// ExtractHorizontalPodAutoscalerStatus is the same as ExtractHorizontalPodAutoscaler except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractHorizontalPodAutoscalerStatus(horizontalPodAutoscaler *apiautoscalingv1.HorizontalPodAutoscaler, fieldManager string) (*HorizontalPodAutoscalerApplyConfiguration, error) {
	return extractHorizontalPodAutoscaler(horizontalPodAutoscaler, fieldManager, "status")
}

func extractHorizontalPodAutoscaler(horizontalPodAutoscaler *apiautoscalingv1.HorizontalPodAutoscaler, fieldManager string, subresource string) (*HorizontalPodAutoscalerApplyConfiguration, error) {
	b := &HorizontalPodAutoscalerApplyConfiguration{}
	err := managedfields.ExtractInto(horizontalPodAutoscaler, internal.Parser().Type("io.k8s.api.autoscaling.v1.HorizontalPodAutoscaler"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(horizontalPodAutoscaler.Name)
	b.WithNamespace(horizontalPodAutoscaler.Namespace)

	b.WithKind("HorizontalPodAutoscaler")
	b.WithAPIVersion("autoscaling/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithKind(value string) *HorizontalPodAutoscalerApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithAPIVersion(value string) *HorizontalPodAutoscalerApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithName(value string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithGenerateName(value string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithNamespace(value string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithSelfLink(value string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.SelfLink = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithUID(value types.UID) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithResourceVersion(value string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithGeneration(value int64) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithCreationTimestamp(value metav1.Time) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithLabels(entries map[string]string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithAnnotations(entries map[string]string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithFinalizers(values ...string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithClusterName(value string) *HorizontalPodAutoscalerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *HorizontalPodAutoscalerApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithSpec(value *HorizontalPodAutoscalerSpecApplyConfiguration) *HorizontalPodAutoscalerApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *HorizontalPodAutoscalerApplyConfiguration) WithStatus(value *HorizontalPodAutoscalerStatusApplyConfiguration) *HorizontalPodAutoscalerApplyConfiguration {
	b.Status = value
	return b
}
