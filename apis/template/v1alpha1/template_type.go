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

type WidgetProperties struct {
	Item `json:",inline"`

	// Key:
	// +optional
	Key string `json:"key,omitempty"`

	// Style:
	// +optional
	Style string `json:"style,omitempty"`

	// Items:
	// +optional
	Items []Item `json:"values,omitempty"`
}

type Widget struct {
	// Title: template title
	Title string `json:"title"`

	// Description: template description
	// +optional
	Description string `json:"description,omitempty"`

	// Properties:
	Properties []WidgetProperties `json:"properties"`
}

type TemplateSpec struct {
	// Title: template title
	Title string `json:"title"`

	// Description: template description
	// +optional
	Description string `json:"description,omitempty"`

	// Icon: template icon
	// +optional
	Icon string `json:"icon,omitempty"`

	// Owner: template owner
	// +optional
	Owner string `json:"owner,omitempty"`
	
	// Url: template file location
	// +optional
	Url string `json:"url,omitempty"`
	
	// EndpointName: endpoint name for this template
	// +optional
	EndpointName string `json:"endpointName,omitempty"`

	// Tags: template tags
	// +optional
	Tags []string `json:"tags,omitempty"`

	// Widgets:
	Widgets []Widget `json:"widgets"`
}

// +kubebuilder:object:root=true

// A Template is a Krateo template API type.
// +kubebuilder:resource:scope=Cluster,categories={managed,krateo,templates}
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TemplateSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// TemplateList contains a list of Template
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TemplateList `json:"items"`
}

// Template type metadata.
var (
	TemplateKind             = reflect.TypeOf(Template{}).Name()
	TemplateGroupKind        = schema.GroupKind{Group: Group, Kind: TemplateKind}.String()
	TemplateKindAPIVersion   = TemplateKind + "." + SchemeGroupVersion.String()
	TemplateGroupVersionKind = SchemeGroupVersion.WithKind(TemplateKind)
)

func init() {
	SchemeBuilder.Register(&Template{}, &TemplateList{})
}
