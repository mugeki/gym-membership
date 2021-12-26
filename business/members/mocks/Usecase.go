// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	members "gym-membership/business/members"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Insert provides a mock function with given fields: classData
func (_m *Repository) Insert(membersData *members.Domain) (members.Domain, error) {
	ret := _m.Called(membersData)

	var r0 members.Domain
	if rf, ok := ret.Get(0).(func(*members.Domain) members.Domain); ok {
		r0 = rf(membersData)
	} else {
		r0 = ret.Get(0).(members.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*members.Domain) error); ok {
		r1 = rf(membersData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateKuota provides a mock function with given fields: idClass
func (_m *Repository) GetByUserID(idmembers int) (string, error) {
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
