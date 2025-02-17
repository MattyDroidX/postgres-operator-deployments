// +kubebuilder:object:generate=true
// +groupName=example.com

package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// GroupVersion is used to register these objects
var (
    GroupVersion = schema.GroupVersion{
        Group:   "example.com",
        Version: "v1",
    }

    // SchemeBuilder is used to add go types to the GroupVersionKind scheme
    SchemeBuilder = scheme.Builder{GroupVersion: GroupVersion}
    AddToScheme   = SchemeBuilder.AddToScheme
)
