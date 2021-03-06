// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import ec "github.com/kyma-incubator/bullseye-showcase/backend/internal/ec"
import mock "github.com/stretchr/testify/mock"

// ProductCacheService is an autogenerated mock type for the ProductCacheService type
type ProductCacheService struct {
	mock.Mock
}

// ForceUpdateProductByID provides a mock function with given fields: ID
func (_m *ProductCacheService) ForceUpdateProductByID(ID string) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForceUpdateProducts provides a mock function with given fields:
func (_m *ProductCacheService) ForceUpdateProducts() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProductDetailsByID provides a mock function with given fields: ID
func (_m *ProductCacheService) GetProductDetailsByID(ID string) (ec.ProductDTO, error) {
	ret := _m.Called(ID)

	var r0 ec.ProductDTO
	if rf, ok := ret.Get(0).(func(string) ec.ProductDTO); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(ec.ProductDTO)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProducts provides a mock function with given fields:
func (_m *ProductCacheService) UpdateProducts() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
