// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	class "gym-membership/business/class"
	class_transactions "gym-membership/business/class_transactions"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetActiveClass provides a mock function with given fields: idUser
func (_m *Usecase) GetActiveClass(idUser uint) ([]class.Domain, error) {
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

// GetAll provides a mock function with given fields: status, idUser, page
func (_m *Usecase) GetAll(status string, idUser uint, page int) ([]class_transactions.Domain, int, int, int64, error) {
	ret := _m.Called(status, idUser, page)

	var r0 []class_transactions.Domain
	if rf, ok := ret.Get(0).(func(string, uint, int) []class_transactions.Domain); ok {
		r0 = rf(status, idUser, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]class_transactions.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string, uint, int) int); ok {
		r1 = rf(status, idUser, page)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(string, uint, int) int); ok {
		r2 = rf(status, idUser, page)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 int64
	if rf, ok := ret.Get(3).(func(string, uint, int) int64); ok {
		r3 = rf(status, idUser, page)
	} else {
		r3 = ret.Get(3).(int64)
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(string, uint, int) error); ok {
		r4 = rf(status, idUser, page)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// GetAllByUser provides a mock function with given fields: idUser
func (_m *Usecase) GetAllByUser(idUser uint) ([]class_transactions.Domain, error) {
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
func (_m *Usecase) Insert(classTransactioData *class_transactions.Domain) (class_transactions.Domain, error) {
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
func (_m *Usecase) UpdateReceipt(idTransactionClass uint, urlImage string) (string, error) {
	ret := _m.Called(idTransactionClass, urlImage)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint, string) string); ok {
		r0 = rf(idTransactionClass, urlImage)
	} else {
		r0 = ret.Get(0).(string)
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
func (_m *Usecase) UpdateStatus(idTransactionClass uint, idAdmin uint, status string) (string, error) {
	ret := _m.Called(idTransactionClass, idAdmin, status)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint, uint, string) string); ok {
		r0 = rf(idTransactionClass, idAdmin, status)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, string) error); ok {
		r1 = rf(idTransactionClass, idAdmin, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
