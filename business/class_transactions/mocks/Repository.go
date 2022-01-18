// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	class "gym-membership/business/class"
	class_transactions "gym-membership/business/class_transactions"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetActiveClass provides a mock function with given fields: idUser
func (_m *Repository) GetActiveClass(idUser uint) ([]class.Domain, error) {
	ret := _m.Called(idUser)

	var r0 []class.Domain
	if rf, ok := ret.Get(0).(func(uint) []class.Domain); ok {
		r0 = rf(idUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]class.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: status, idUser, offset, limit
func (_m *Repository) GetAll(status string, idUser uint, offset int, limit int) ([]class_transactions.Domain, int64, error) {
	ret := _m.Called(status, idUser, offset, limit)

	var r0 []class_transactions.Domain
	if rf, ok := ret.Get(0).(func(string, uint, int, int) []class_transactions.Domain); ok {
		r0 = rf(status, idUser, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]class_transactions.Domain)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(string, uint, int, int) int64); ok {
		r1 = rf(status, idUser, offset, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, uint, int, int) error); ok {
		r2 = rf(status, idUser, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllByUser provides a mock function with given fields: idUser
func (_m *Repository) GetAllByUser(idUser uint) ([]class_transactions.Domain, error) {
	ret := _m.Called(idUser)

	var r0 []class_transactions.Domain
	if rf, ok := ret.Get(0).(func(uint) []class_transactions.Domain); ok {
		r0 = rf(idUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]class_transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: classTransactioData
func (_m *Repository) Insert(classTransactioData *class_transactions.Domain) (class_transactions.Domain, error) {
	ret := _m.Called(classTransactioData)

	var r0 class_transactions.Domain
	if rf, ok := ret.Get(0).(func(*class_transactions.Domain) class_transactions.Domain); ok {
		r0 = rf(classTransactioData)
	} else {
		r0 = ret.Get(0).(class_transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*class_transactions.Domain) error); ok {
		r1 = rf(classTransactioData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReceipt provides a mock function with given fields: idTransactionClass, urlImage
func (_m *Repository) UpdateReceipt(idTransactionClass uint, urlImage string) (class_transactions.Domain, error) {
	ret := _m.Called(idTransactionClass, urlImage)

	var r0 class_transactions.Domain
	if rf, ok := ret.Get(0).(func(uint, string) class_transactions.Domain); ok {
		r0 = rf(idTransactionClass, urlImage)
	} else {
		r0 = ret.Get(0).(class_transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(idTransactionClass, urlImage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: idTransactionClass, idAdmin, status
func (_m *Repository) UpdateStatus(idTransactionClass uint, idAdmin uint, status string) (class_transactions.Domain, error) {
	ret := _m.Called(idTransactionClass, idAdmin, status)

	var r0 class_transactions.Domain
	if rf, ok := ret.Get(0).(func(uint, uint, string) class_transactions.Domain); ok {
		r0 = rf(idTransactionClass, idAdmin, status)
	} else {
		r0 = ret.Get(0).(class_transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, string) error); ok {
		r1 = rf(idTransactionClass, idAdmin, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
