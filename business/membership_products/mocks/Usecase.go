// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	membership_products "gym-membership/business/membership_products"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: id
func (_m *Usecase) DeleteByID(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: page
func (_m *Usecase) GetAll(page int) ([]membership_products.Domain, int, int, int64, error) {
	ret := _m.Called(page)

	var r0 []membership_products.Domain
	if rf, ok := ret.Get(0).(func(int) []membership_products.Domain); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]membership_products.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(int) int); ok {
		r2 = rf(page)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 int64
	if rf, ok := ret.Get(3).(func(int) int64); ok {
		r3 = rf(page)
	} else {
		r3 = ret.Get(3).(int64)
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(int) error); ok {
		r4 = rf(page)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// GetByID provides a mock function with given fields: id
func (_m *Usecase) GetByID(id uint) (membership_products.Domain, error) {
	ret := _m.Called(id)

	var r0 membership_products.Domain
	if rf, ok := ret.Get(0).(func(uint) membership_products.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(membership_products.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newData
func (_m *Usecase) Insert(newData *membership_products.Domain) (membership_products.Domain, error) {
	ret := _m.Called(newData)

	var r0 membership_products.Domain
	if rf, ok := ret.Get(0).(func(*membership_products.Domain) membership_products.Domain); ok {
		r0 = rf(newData)
	} else {
		r0 = ret.Get(0).(membership_products.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*membership_products.Domain) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: id, newData
func (_m *Usecase) UpdateByID(id uint, newData *membership_products.Domain) (membership_products.Domain, error) {
	ret := _m.Called(id, newData)

	var r0 membership_products.Domain
	if rf, ok := ret.Get(0).(func(uint, *membership_products.Domain) membership_products.Domain); ok {
		r0 = rf(id, newData)
	} else {
		r0 = ret.Get(0).(membership_products.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *membership_products.Domain) error); ok {
		r1 = rf(id, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
