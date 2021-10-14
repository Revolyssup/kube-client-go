package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GV = schema.GroupVersion{
	Group:   "goglot.dev",
	Version: "v1alpha1",
}

var (
	SchemeBuilder runtime.SchemeBuilder
	AddToScheme   = SchemeBuilder.AddToScheme //needed by code generator
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GV, &GlotPod{}, &GlotPodList{})
	metav1.AddToGroupVersion(scheme, GV)
	return nil
}

func init() {
	SchemeBuilder.Register(addKnownTypes)
}
