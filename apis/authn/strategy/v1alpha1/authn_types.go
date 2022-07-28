package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// StrategyParams describe an AuthN Krateo strategy.
type StrategyParams struct {
	// Name: descriptive name for this strategy.
	Name string `json:"name"`

	// Config: authn configuration data.
	Config string `json:"config"`

	// Strategy: type of AuthN strategy (i.e. basic, github, microsoft, etc.).
	Strategy string `json:"strategy"`

	// Type: type of this strategy handler.
	Type string `json:"type"`
}

// +kubebuilder:object:root=true

// A Strategy is a AuthN API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status",priority=1
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status",priority=1
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={managed,krateo,authn}
type Strategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec StrategyParams `json:"spec"`
}

// +kubebuilder:object:root=true

// StrategyList contains a list of Strategy
type StrategyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StrategyList `json:"items"`
}

// Strategy type metadata.
var (
	StrategyKind             = reflect.TypeOf(Strategy{}).Name()
	StrategyGroupKind        = schema.GroupKind{Group: Group, Kind: StrategyKind}.String()
	StrategyKindAPIVersion   = StrategyKind + "." + SchemeGroupVersion.String()
	StrategyGroupVersionKind = SchemeGroupVersion.WithKind(StrategyKind)
)

func init() {
	SchemeBuilder.Register(&Strategy{}, &StrategyList{})
}
