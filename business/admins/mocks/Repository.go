// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	admins "gym-membership/business/admins"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: id
func (_m *Repository) DeleteByID(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: name, offset, limit
func (_m *Repository) GetAll(name string, offset int, limit int) ([]admins.Domain, int64, error) {
	ret := _m.Called(name, offset, limit)

	var r0 []admins.Domain
	if rf, ok := ret.Get(0).(func(string, int, int) []admins.Domain); ok {
		r0 = rf(name, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]admins.Domain)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(string, int, int) int64); ok {
		r1 = rf(name, offset, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int, int) error); ok {
		r2 = rf(name, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id uint) (admins.Domain, error) {
	ret := _m.Called(id)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(uint) admins.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *Repository) GetByUsername(username string) (admins.Domain, error) {
	ret := _m.Called(username)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(string) admins.Domain); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: adminData
func (_m *Repository) Register(adminData *admins.Domain) (admins.Domain, error) {
	ret := _m.Called(adminData)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(*admins.Domain) admins.Domain); ok {
		r0 = rf(adminData)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*admins.Domain) error); ok {
		r1 = rf(adminData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, adminData
func (_m *Repository) Update(id uint, adminData *admins.Domain) (admins.Domain, error) {
	ret := _m.Called(id, adminData)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(uint, *admins.Domain) admins.Domain); ok {
		r0 = rf(id, adminData)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *admins.Domain) error); ok {
		r1 = rf(id, adminData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
