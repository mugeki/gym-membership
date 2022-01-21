// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	membership_transactions "gym-membership/business/membership_transactions"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: date, status, idUser, offset, limit
func (_m *Repository) GetAll(date time.Time, status string, idUser uint, offset int, limit int) ([]membership_transactions.Domain, int64, error) {
	ret := _m.Called(date, status, idUser, offset, limit)

	var r0 []membership_transactions.Domain
	if rf, ok := ret.Get(0).(func(time.Time, string, uint, int, int) []membership_transactions.Domain); ok {
		r0 = rf(date, status, idUser, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]membership_transactions.Domain)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(time.Time, string, uint, int, int) int64); ok {
		r1 = rf(date, status, idUser, offset, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(time.Time, string, uint, int, int) error); ok {
		r2 = rf(date, status, idUser, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllByUser provides a mock function with given fields: idUser
func (_m *Repository) GetAllByUser(idUser uint) ([]membership_transactions.Domain, error) {
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
func (_m *Repository) GetByID(idTransaction uint) (membership_transactions.Domain, error) {
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
func (_m *Repository) Insert(membershipTransactionData *membership_transactions.Domain) (membership_transactions.Domain, error) {
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
func (_m *Repository) UpdateReceipt(idTransactionClass uint, urlImage string) (membership_transactions.Domain, error) {
	ret := _m.Called(idTransactionClass, urlImage)

	var r0 membership_transactions.Domain
	if rf, ok := ret.Get(0).(func(uint, string) membership_transactions.Domain); ok {
		r0 = rf(idTransactionClass, urlImage)
	} else {
		r0 = ret.Get(0).(membership_transactions.Domain)
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
func (_m *Repository) UpdateStatus(id uint, idAdmin uint, status string) (membership_transactions.Domain, error) {
	ret := _m.Called(id, idAdmin, status)

	var r0 membership_transactions.Domain
	if rf, ok := ret.Get(0).(func(uint, uint, string) membership_transactions.Domain); ok {
		r0 = rf(id, idAdmin, status)
	} else {
		r0 = ret.Get(0).(membership_transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, string) error); ok {
		r1 = rf(id, idAdmin, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
