package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Item struct {
	// Title:
	// +optional
	Title string `json:"title"`

	// Type:
	// +optional
	Type string `json:"type"`

	// Description:
	// +optional
	Description string `json:"description,omitempty"`

	// Value:
	// +optional
	Value string `json:"value,omitempty"`

	// Default:
	// +optional
	Default string `json:"default,omitempty"`

	// Cost:
	// +optional
	Cost string `json:"cost,omitempty"`

	// Required:
	Required *bool `json:"required,omitempty"`

	// Visible:
	// +optional
	Visible *bool `json:"visible,omitempty"`

	// ReadOnly:
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// Link to Krateo deployment plugins
type Link struct {
	// Title: link title
	Title string `json:"title"`

	// Value: link value (url)
	Value string `json:"value"`

	// Type: link type
	Type string `json:"type"`

	// Icon: link icon
	// +optional
	Icon string `json:"icon,omitempty"`
}

type Plugin struct {
	// Title: link title
	Title string `json:"title"`

	// Value: link value (url)
	// +optional
	Value string `json:"value"`

	// Type: link type
	Type string `json:"type"`

	// Icon: link icon
	// +optional
	Icon string `json:"icon,omitempty"`

	// EndpointName:
	// +optional
	EndpointName string `json:"endpointName,omitempty"`

	// Values:
	// +optional
	Values []string `json:"values,omitempty"`
}

type DeploymentSpec struct {
	// Title: deployment title
	Title string `json:"title"`

	// Description: deployment description
	// +optional
	Description string `json:"description,omitempty"`

	// Icon: deployment icon
	// +optional
	Icon string `json:"icon,omitempty"`

	// Owner: deployment owner
	// +optional
	Owner string `json:"owner,omitempty"`
	
	// Url: deployment file location
	// +optional
	Url string `json:"url,omitempty"`
	
	// EndpointName: endpoint name for this deployment
	// +optional
	EndpointName string `json:"endpointName,omitempty"`

	// Tags: deployment tags
	// +optional
	Tags []string `json:"tags,omitempty"`

	// Links: deployment plugins
	// +optional
	Links []Link `json:"links,omitempty"`

	// Plugins: deployment plugins
	// +optional
	Plugins []Plugin `json:"plugins,omitempty"`
}

// +kubebuilder:object:root=true

// A Deployment is a Krateo deployment API type.
// +kubebuilder:resource:scope=Cluster,categories={managed,krateo,deployments}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DeploymentSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployment
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentList `json:"items"`
}

// Deployment type metadata.
var (
	DeploymentKind             = reflect.TypeOf(Deployment{}).Name()
	DeploymentGroupKind        = schema.GroupKind{Group: Group, Kind: DeploymentKind}.String()
	DeploymentKindAPIVersion   = DeploymentKind + "." + SchemeGroupVersion.String()
	DeploymentGroupVersionKind = SchemeGroupVersion.WithKind(DeploymentKind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
