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
	apipolicyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "github.com/Angus-F/client-go/applyconfigurations/internal"
	v1 "github.com/Angus-F/client-go/applyconfigurations/meta/v1"
)

// PodDisruptionBudgetApplyConfiguration represents an declarative configuration of the PodDisruptionBudget type for use
// with apply.
type PodDisruptionBudgetApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *PodDisruptionBudgetSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *PodDisruptionBudgetStatusApplyConfiguration `json:"status,omitempty"`
}

// PodDisruptionBudget constructs an declarative configuration of the PodDisruptionBudget type for use with
// apply.
func PodDisruptionBudget(name, namespace string) *PodDisruptionBudgetApplyConfiguration {
	b := &PodDisruptionBudgetApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("PodDisruptionBudget")
	b.WithAPIVersion("policy/v1")
	return b
}

// ExtractPodDisruptionBudget extracts the applied configuration owned by fieldManager from
// podDisruptionBudget. If no managedFields are found in podDisruptionBudget for fieldManager, a
// PodDisruptionBudgetApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. Is is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// podDisruptionBudget must be a unmodified PodDisruptionBudget API object that was retrieved from the Kubernetes API.
// ExtractPodDisruptionBudget provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractPodDisruptionBudget(podDisruptionBudget *apipolicyv1.PodDisruptionBudget, fieldManager string) (*PodDisruptionBudgetApplyConfiguration, error) {
	return extractPodDisruptionBudget(podDisruptionBudget, fieldManager, "")
}

// ExtractPodDisruptionBudgetStatus is the same as ExtractPodDisruptionBudget except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractPodDisruptionBudgetStatus(podDisruptionBudget *apipolicyv1.PodDisruptionBudget, fieldManager string) (*PodDisruptionBudgetApplyConfiguration, error) {
	return extractPodDisruptionBudget(podDisruptionBudget, fieldManager, "status")
}

func extractPodDisruptionBudget(podDisruptionBudget *apipolicyv1.PodDisruptionBudget, fieldManager string, subresource string) (*PodDisruptionBudgetApplyConfiguration, error) {
	b := &PodDisruptionBudgetApplyConfiguration{}
	err := managedfields.ExtractInto(podDisruptionBudget, internal.Parser().Type("io.k8s.api.policy.v1.PodDisruptionBudget"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(podDisruptionBudget.Name)
	b.WithNamespace(podDisruptionBudget.Namespace)

	b.WithKind("PodDisruptionBudget")
	b.WithAPIVersion("policy/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithKind(value string) *PodDisruptionBudgetApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithAPIVersion(value string) *PodDisruptionBudgetApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithName(value string) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithGenerateName(value string) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithNamespace(value string) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithSelfLink(value string) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.SelfLink = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithUID(value types.UID) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithResourceVersion(value string) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithGeneration(value int64) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithCreationTimestamp(value metav1.Time) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *PodDisruptionBudgetApplyConfiguration) WithLabels(entries map[string]string) *PodDisruptionBudgetApplyConfiguration {
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
func (b *PodDisruptionBudgetApplyConfiguration) WithAnnotations(entries map[string]string) *PodDisruptionBudgetApplyConfiguration {
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
func (b *PodDisruptionBudgetApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *PodDisruptionBudgetApplyConfiguration {
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
func (b *PodDisruptionBudgetApplyConfiguration) WithFinalizers(values ...string) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithClusterName(value string) *PodDisruptionBudgetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *PodDisruptionBudgetApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithSpec(value *PodDisruptionBudgetSpecApplyConfiguration) *PodDisruptionBudgetApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *PodDisruptionBudgetApplyConfiguration) WithStatus(value *PodDisruptionBudgetStatusApplyConfiguration) *PodDisruptionBudgetApplyConfiguration {
	b.Status = value
	return b
}
