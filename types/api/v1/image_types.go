// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"github.com/opencontainers/go-digest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Layer struct {
	ID    digest.Digest `json:"id,omitempty"` // shaxxx:d6a6c9bfd4ad2901695be1dceca62e1c35a8482982ad6be172fe6958bc4f79d7
	Type  string        `json:"type,omitempty"`
	Value string        `json:"value,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Image. Edit Image_types.go to remove/update
	ID            string   `json:"id,omitempty"`
	MergedLayer   string   `json:"mergedLayer,omitempty"`
	Layers        []Layer  `json:"layers,omitempty"`
	SealerVersion string   `json:"sealer_version,omitempty"`
	Platform      Platform `json:"platform"`
}

// ImageStatus defines the observed state of Image
type ImageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Image is the Schema for the images API
type Image struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   ImageSpec   `json:"spec,omitempty"  yaml:"spec,omitempty"`
	Status ImageStatus `json:"status,omitempty"  yaml:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Image
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items,omitempty"`
}

type Platform struct {
	Architecture string `json:"architecture,omitempty"`
	OS           string `json:"os,omitempty"`
	OSVersion    string `json:"os_version,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
