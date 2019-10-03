package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BattlefieldSpec defines the desired state of Battlefield
// +k8s:openapi-gen=true
type BattlefieldSpec struct {
	Duration     int      `json:"duration"`
	Players      []Player `json:"players"`
	HitFrequency int      `json:"hitFrequency"`

	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// BattlefieldStatus defines the observed state of Battlefield
// +k8s:openapi-gen=true
type BattlefieldStatus struct {
	Phase     string         `json:"phase"`
	StartTime *metav1.Time   `json:"startTime,omitempty"`
	StopTime  *metav1.Time   `json:"stopTime,omitempty"`
	Scores    []PlayerStatus `json:"scores"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Battlefield is the Schema for the battlefields API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Battlefield struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BattlefieldSpec   `json:"spec,omitempty"`
	Status BattlefieldStatus `json:"status,omitempty"`
}

// Player defines one container player
type Player struct {
	Name      string `json:"name"`
	Image     string `json:"image,omitempty"`
	MaxHealth int    `json:"maxhealth"`
	Shield 	  bool   `json:"shield"`
	Disqualified 	  bool   `json:"disqualified"`
}

// PlayerStatus records score of a player
type PlayerStatus struct {
	Name  string `json:"name"`
	Kill  int    `json:"kill"`
	Death int    `json:"death"`
	Ready bool   `json:"ready"`
	KilledBy string 	`json:"killedby"`
	CurrentHealth int	`json:"currentHealth"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BattlefieldList contains a list of Battlefield
type BattlefieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Battlefield `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Battlefield{}, &BattlefieldList{})
}
