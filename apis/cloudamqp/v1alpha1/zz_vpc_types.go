// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPCInitParameters struct {

	// The name of the VPC.
	// Name of the VPC instance
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The hosted region for the managed standalone VPC
	// The hosted region for the standalone VPC instance
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The VPC subnet
	// The VPC subnet
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// Tag the VPC with optional tags
	// Tag the VPC instance with optional tags
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPCObservation struct {

	// The identifier for this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the VPC.
	// Name of the VPC instance
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The hosted region for the managed standalone VPC
	// The hosted region for the standalone VPC instance
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The VPC subnet
	// The VPC subnet
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// Tag the VPC with optional tags
	// Tag the VPC instance with optional tags
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// VPC name given when hosted at the cloud provider
	// VPC name given when hosted at the cloud provider
	VPCName *string `json:"vpcName,omitempty" tf:"vpc_name,omitempty"`
}

type VPCParameters struct {

	// The name of the VPC.
	// Name of the VPC instance
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The hosted region for the managed standalone VPC
	// The hosted region for the standalone VPC instance
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The VPC subnet
	// The VPC subnet
	// +kubebuilder:validation:Optional
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// Tag the VPC with optional tags
	// Tag the VPC instance with optional tags
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCSpec defines the desired state of VPC
type VPCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VPCInitParameters `json:"initProvider,omitempty"`
}

// VPCStatus defines the observed state of VPC.
type VPCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPC is the Schema for the VPCs API. Managed VPC resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudamqp}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnet) || (has(self.initProvider) && has(self.initProvider.subnet))",message="spec.forProvider.subnet is a required parameter"
	Spec   VPCSpec   `json:"spec"`
	Status VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPCs
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

// Repository type metadata.
var (
	VPC_Kind             = "VPC"
	VPC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPC_Kind}.String()
	VPC_KindAPIVersion   = VPC_Kind + "." + CRDGroupVersion.String()
	VPC_GroupVersionKind = CRDGroupVersion.WithKind(VPC_Kind)
)

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}
