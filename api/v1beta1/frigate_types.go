/*
Copyright 2020 Ace Eldeib.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FrigateSpec defines the desired state of Frigate
type FrigateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Frigate. Edit Frigate_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// FrigateStatus defines the observed state of Frigate
type FrigateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Frigate is the Schema for the frigates API
type Frigate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FrigateSpec   `json:"spec,omitempty"`
	Status FrigateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrigateList contains a list of Frigate
type FrigateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Frigate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Frigate{}, &FrigateList{})
}
