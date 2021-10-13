package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// These const variables are used in our custom controller.
const (
	GroupName string = "mycontroller.tatacommunications.com"
	Kind      string = "TestResource"
	Version   string = "v1alpha1"
	// Plural    string = "testresources"
	// Singluar  string = "testresource"
	// ShortName string = "ts"
	// Name      string = Plural + "." + GroupName
)

// TestResourceSpec specifies the 'spec' of TestResource CRD.
type TestResourceSpec struct {
	FirstNum  *int32 `json:"firstNum"`
	SecondNum *int32 `json:"secondNum"`
	Operation string `json:"operation"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestResource describes a TestResource custom resource.
type TestResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TestResourceSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestResourceList is a list of TestResource resources.
type TestResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TestResource `json:"items"`
}
