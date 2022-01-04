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

// GetByUserID provides a mock function with given fields: idMembers
func (_m *Usecase) GetByUserID(idMembers int) (string, error) {
	ret := _m.Called(idMembers)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(idMembers)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idMembers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: membersData
func (_m *Usecase) Insert(membersData *membership_products.Domain) (string, error) {
	ret := _m.Called(membersData)

	var r0 string
	if rf, ok := ret.Get(0).(func(*membership_products.Domain) string); ok {
		r0 = rf(membersData)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*membership_products.Domain) error); ok {
		r1 = rf(membersData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Usecase) UpdateStatus(idMembers uint) (string, error) {
	ret := _m.Called(idMembers)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint) string); ok {
		r0 = rf(idMembers)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idMembers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Usecase) DeleteByID(idMembers uint) error {
	ret := _m.Called(idMembers)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(idMembers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Usecase) UpdateByID(idMembers uint, membersData *membership_products.Domain) (membership_products.Domain, error) {
	ret := _m.Called(idMembers, membersData)

	var r0 membership_products.Domain
	if rf, ok := ret.Get(0).(func(uint, *membership_products.Domain) membership_products.Domain); ok {
		r0 = rf(idMembers, membersData)
	} else {
		r0 = ret.Get(0).(membership_products.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *membership_products.Domain) error); ok {
		r1 = rf(idMembers, membersData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
