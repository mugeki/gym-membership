// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	payment_accounts "gym-membership/business/payment_accounts"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *Usecase) GetAll() ([]payment_accounts.Domain, error) {
	ret := _m.Called()

	var r0 []payment_accounts.Domain
	if rf, ok := ret.Get(0).(func() []payment_accounts.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]payment_accounts.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: account
func (_m *Usecase) Insert(account *payment_accounts.Domain) (payment_accounts.Domain, error) {
	ret := _m.Called(account)

	var r0 payment_accounts.Domain
	if rf, ok := ret.Get(0).(func(*payment_accounts.Domain) payment_accounts.Domain); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Get(0).(payment_accounts.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*payment_accounts.Domain) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
