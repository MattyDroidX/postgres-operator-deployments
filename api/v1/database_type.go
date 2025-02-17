package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=databases,scope=Namespaced,shortName=db
// +kubebuilder:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Database defines a PostgreSQL Database custom resource
type Database struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    Spec              DatabaseSpec   `json:"spec,omitempty"`
    Status            DatabaseStatus `json:"status,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
    Size    string `json:"size,omitempty"`    // small, medium, large
    Version string `json:"version,omitempty"` // e.g. "14"
}

// DatabaseStatus defines the observed state of Database
type DatabaseStatus struct {
    Status string `json:"status,omitempty"` // e.g. "Provisioned"
}

// +kubebuilder:object:root=true
// +kubebuilder:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatabaseList contains a list of Database items
type DatabaseList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []Database `json:"items"`
}

func init() {
    // The actual registration is done via groupversion_info.go (SchemeBuilder)
}
