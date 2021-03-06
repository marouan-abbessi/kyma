// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"

import v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"

// ServiceBindingFinderLister is an autogenerated mock type for the ServiceBindingFinderLister type
type ServiceBindingFinderLister struct {
	mock.Mock
}

// Find provides a mock function with given fields: env, name
func (_m *ServiceBindingFinderLister) Find(env string, name string) (*v1beta1.ServiceBinding, error) {
	ret := _m.Called(env, name)

	var r0 *v1beta1.ServiceBinding
	if rf, ok := ret.Get(0).(func(string, string) *v1beta1.ServiceBinding); ok {
		r0 = rf(env, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ServiceBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(env, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForServiceInstance provides a mock function with given fields: env, instanceName
func (_m *ServiceBindingFinderLister) ListForServiceInstance(env string, instanceName string) ([]*v1beta1.ServiceBinding, error) {
	ret := _m.Called(env, instanceName)

	var r0 []*v1beta1.ServiceBinding
	if rf, ok := ret.Get(0).(func(string, string) []*v1beta1.ServiceBinding); ok {
		r0 = rf(env, instanceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.ServiceBinding)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(env, instanceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
