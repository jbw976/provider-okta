/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BehaviorObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Determines the method and level of detail used to evaluate the behavior.
	LocationGranularityType *string `json:"locationGranularityType,omitempty" tf:"location_granularity_type,omitempty"`

	// Name of the behavior
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of recent authentications used to evaluate the behavior.
	NumberOfAuthentications *float64 `json:"numberOfAuthentications,omitempty" tf:"number_of_authentications,omitempty"`

	// Radius from location (in kilometers)
	RadiusFromLocation *float64 `json:"radiusFromLocation,omitempty" tf:"radius_from_location,omitempty"`

	// Behavior status: ACTIVE or INACTIVE.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Behavior type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Velocity (in kilometers per hour).
	Velocity *float64 `json:"velocity,omitempty" tf:"velocity,omitempty"`
}

type BehaviorParameters struct {

	// Determines the method and level of detail used to evaluate the behavior.
	// +kubebuilder:validation:Optional
	LocationGranularityType *string `json:"locationGranularityType,omitempty" tf:"location_granularity_type,omitempty"`

	// Name of the behavior
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of recent authentications used to evaluate the behavior.
	// +kubebuilder:validation:Optional
	NumberOfAuthentications *float64 `json:"numberOfAuthentications,omitempty" tf:"number_of_authentications,omitempty"`

	// Radius from location (in kilometers)
	// +kubebuilder:validation:Optional
	RadiusFromLocation *float64 `json:"radiusFromLocation,omitempty" tf:"radius_from_location,omitempty"`

	// Behavior status: ACTIVE or INACTIVE.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Behavior type
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Velocity (in kilometers per hour).
	// +kubebuilder:validation:Optional
	Velocity *float64 `json:"velocity,omitempty" tf:"velocity,omitempty"`
}

// BehaviorSpec defines the desired state of Behavior
type BehaviorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BehaviorParameters `json:"forProvider"`
}

// BehaviorStatus defines the observed state of Behavior.
type BehaviorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BehaviorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Behavior is the Schema for the Behaviors API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,okta}
type Behavior struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.type)",message="type is a required parameter"
	Spec   BehaviorSpec   `json:"spec"`
	Status BehaviorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BehaviorList contains a list of Behaviors
type BehaviorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Behavior `json:"items"`
}

// Repository type metadata.
var (
	Behavior_Kind             = "Behavior"
	Behavior_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Behavior_Kind}.String()
	Behavior_KindAPIVersion   = Behavior_Kind + "." + CRDGroupVersion.String()
	Behavior_GroupVersionKind = CRDGroupVersion.WithKind(Behavior_Kind)
)

func init() {
	SchemeBuilder.Register(&Behavior{}, &BehaviorList{})
}