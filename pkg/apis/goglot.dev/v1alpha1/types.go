package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GlotPodSpec struct {
	Name     string
	Language string
	Code     string
	Input    string
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GlotPod struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec GlotPodSpec
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GlotPodList struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Items []GlotPod
}
