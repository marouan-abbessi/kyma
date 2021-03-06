// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import v1alpha1 "github.com/kyma-project/kyma/components/application-operator/pkg/apis/applicationconnector/v1alpha1"

// ApplicationLister is an autogenerated failing mock type for the ApplicationLister type
type ApplicationLister struct {
	err error
}

// NewApplicationLister creates a new ApplicationLister type instance
func NewApplicationLister(err error) *ApplicationLister {
	return &ApplicationLister{err: err}
}

// ListInEnvironment provides a failing mock function with given fields: environment
func (_m *ApplicationLister) ListInEnvironment(environment string) ([]*v1alpha1.Application, error) {
	var r0 []*v1alpha1.Application
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ListNamespacesFor provides a failing mock function with given fields: appName
func (_m *ApplicationLister) ListNamespacesFor(appName string) ([]string, error) {
	var r0 []string
	var r1 error
	r1 = _m.err

	return r0, r1
}
