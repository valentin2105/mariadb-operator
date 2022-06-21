/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RestoreMariaDBSpec defines the desired state of RestoreMariaDB
type RestoreMariaDBSpec struct {
	// +kubebuilder:validation:Required
	MariaDBRef corev1.LocalObjectReference `json:"mariaDbRef"`
	// +kubebuilder:validation:Required
	BackupMariaDBRef corev1.LocalObjectReference `json:"backupMariaDBRef"`
	// +kubebuilder:default=3
	BackoffLimit int32 `json:"backoffLimit,omitempty"`
	// +kubebuilder:default=OnFailure
	RestartPolicy corev1.RestartPolicy `json:"restartPolicy,omitempty"`

	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// RestoreMariaDBStatus defines the observed state of RestoreMariaDB
type RestoreMariaDBStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func (r *RestoreMariaDBStatus) AddCondition(condition metav1.Condition) {
	if r.Conditions == nil {
		r.Conditions = make([]metav1.Condition, 0)
	}
	meta.SetStatusCondition(&r.Conditions, condition)
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RestoreMariaDB is the Schema for the restoremariadbs API
type RestoreMariaDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreMariaDBSpec   `json:"spec,omitempty"`
	Status RestoreMariaDBStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RestoreMariaDBList contains a list of RestoreMariaDB
type RestoreMariaDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestoreMariaDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RestoreMariaDB{}, &RestoreMariaDBList{})
}