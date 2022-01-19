// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	membership_transactions "gym-membership/business/membership_transactions"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: status, idUser, page
func (_m *Usecase) GetAll(status string, idUser uint, page int) ([]membership_transactions.Domain, int, int, int64, error) {
	ret := _m.Called(status, idUser, page)

	var r0 []membership_transactions.Domain
	if rf, ok := ret.Get(0).(func(string, uint, int) []membership_transactions.Domain); ok {
		r0 = rf(status, idUser, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]membership_transactions.Domain)
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
func (_m *Usecase) GetAllByUser(idUser uint) ([]membership_transactions.Domain, error) {
	ret := _m.Called(idUser)

	var r0 []membership_transactions.Domain
	if rf, ok := ret.Get(0).(func(uint) []membership_transactions.Domain); ok {
		r0 = rf(idUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]membership_transactions.Domain)
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

// GetByID provides a mock function with given fields: idTransaction
func (_m *Usecase) GetByID(idTransaction uint) (membership_transactions.Domain, error) {
	ret := _m.Called(idTransaction)

	var r0 membership_transactions.Domain
	if rf, ok := ret.Get(0).(func(uint) membership_transactions.Domain); ok {
		r0 = rf(idTransaction)
	} else {
		r0 = ret.Get(0).(membership_transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(idTransaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: membershipTransactionData
func (_m *Usecase) Insert(membershipTransactionData *membership_transactions.Domain) (membership_transactions.Domain, error) {
	ret := _m.Called(membershipTransactionData)

	var r0 membership_transactions.Domain
	if rf, ok := ret.Get(0).(func(*membership_transactions.Domain) membership_transactions.Domain); ok {
		r0 = rf(membershipTransactionData)
	} else {
		r0 = ret.Get(0).(membership_transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*membership_transactions.Domain) error); ok {
		r1 = rf(membershipTransactionData)
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

// UpdateStatus provides a mock function with given fields: id, idAdmin, status
func (_m *Usecase) UpdateStatus(id uint, idAdmin uint, status string) error {
	ret := _m.Called(id, idAdmin, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, string) error); ok {
		r0 = rf(id, idAdmin, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
