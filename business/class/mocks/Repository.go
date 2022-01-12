// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	class "gym-membership/business/class"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: title, offset, limit
func (_m *Repository) GetAll(title string, offset int, limit int) ([]class.Domain, int64, error) {
	ret := _m.Called(title, offset, limit)

	var r0 []class.Domain
	if rf, ok := ret.Get(0).(func(string, int, int) []class.Domain); ok {
		r0 = rf(title, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]class.Domain)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(string, int, int) int64); ok {
		r1 = rf(title, offset, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int, int) error); ok {
		r2 = rf(title, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Insert provides a mock function with given fields: classData
func (_m *Repository) Insert(classData *class.Domain) (class.Domain, error) {
	ret := _m.Called(classData)

	var r0 class.Domain
	if rf, ok := ret.Get(0).(func(*class.Domain) class.Domain); ok {
		r0 = rf(classData)
	} else {
		r0 = ret.Get(0).(class.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*class.Domain) error); ok {
		r1 = rf(classData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsExist provides a mock function with given fields: idClass
func (_m *Repository) IsExist(idClass int) (class.Domain, error) {
	ret := _m.Called(idClass)

	var r0 class.Domain
	if rf, ok := ret.Get(0).(func(int) class.Domain); ok {
		r0 = rf(idClass)
	} else {
		r0 = ret.Get(0).(class.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idClass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClassByID provides a mock function with given fields: id, articleData
func (_m *Repository) UpdateClassByID(id uint, articleData *class.Domain) (class.Domain, error) {
	ret := _m.Called(id, articleData)

	var r0 class.Domain
	if rf, ok := ret.Get(0).(func(uint, *class.Domain) class.Domain); ok {
		r0 = rf(id, articleData)
	} else {
		r0 = ret.Get(0).(class.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *class.Domain) error); ok {
		r1 = rf(id, articleData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateParticipant provides a mock function with given fields: idClass
func (_m *Repository) UpdateParticipant(idClass int) (class.Domain, error) {
	ret := _m.Called(idClass)

	var r0 class.Domain
	if rf, ok := ret.Get(0).(func(int) class.Domain); ok {
		r0 = rf(idClass)
	} else {
		r0 = ret.Get(0).(class.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idClass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: idClass, status
func (_m *Repository) UpdateStatus(idClass int, status bool) (class.Domain, error) {
	ret := _m.Called(idClass, status)

	var r0 class.Domain
	if rf, ok := ret.Get(0).(func(int, bool) class.Domain); ok {
		r0 = rf(idClass, status)
	} else {
		r0 = ret.Get(0).(class.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, bool) error); ok {
		r1 = rf(idClass, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
