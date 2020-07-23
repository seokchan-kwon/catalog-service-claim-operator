package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CatalogServiceClaimSpec defines the desired state of CatalogServiceClaim
type CatalogServiceClaimSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// CatalogServiceClaimStatus defines the observed state of CatalogServiceClaim
type CatalogServiceClaimStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CatalogServiceClaim is the Schema for the catalogserviceclaims API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=catalogserviceclaims,scope=Namespaced
type CatalogServiceClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CatalogServiceClaimSpec   `json:"spec,omitempty"`
	Status CatalogServiceClaimStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CatalogServiceClaimList contains a list of CatalogServiceClaim
type CatalogServiceClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogServiceClaim `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CatalogServiceClaim{}, &CatalogServiceClaimList{})
}
