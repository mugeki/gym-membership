// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	class "gym-membership/business/class"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Insert provides a mock function with given fields: classData
func (_m *Usecase) Insert(classData *class.Domain) (string, error) {
	ret := _m.Called(classData)

	var r0 string
	if rf, ok := ret.Get(0).(func(*class.Domain) string); ok {
		r0 = rf(classData)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*class.Domain) error); ok {
		r1 = rf(classData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateKuota provides a mock function with given fields: idClass
func (_m *Usecase) UpdateKuota(idClass int) (string, error) {
	ret := _m.Called(idClass)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(idClass)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idClass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}