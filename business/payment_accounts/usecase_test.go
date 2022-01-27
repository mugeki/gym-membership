package payment_accounts_test

import (
	"gym-membership/business/payment_accounts"
	_paymentAccountsMock "gym-membership/business/payment_accounts/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockPaymentAccountRepo _paymentAccountsMock.Repository
	paymentAccountsUsecase  payment_accounts.Usecase
	paymentAccountsData     payment_accounts.Domain
)

func TestMain(m *testing.M) {
	paymentAccountsUsecase = payment_accounts.NewPaymentAccountUsecase(&mockPaymentAccountRepo)
	paymentAccountsData = payment_accounts.Domain{
		ID:			1,
		Name:		"BCA",
		NoCard:		"1110002223344",
		OwnerName:	"Test Name",
		Desc:		"A payment account",
	}
	m.Run()
}

func TestGetAll(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockPaymentAccountRepo.On("GetAll").
			Return([]payment_accounts.Domain{paymentAccountsData}, nil).Once()

		resp, err := paymentAccountsUsecase.GetAll()

		assert.Nil(t, err)
		assert.Contains(t, resp, paymentAccountsData)
		assert.Len(t, resp, 1)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockPaymentAccountRepo.On("GetAll").
			Return([]payment_accounts.Domain{}, assert.AnError).Once()

		resp, err := paymentAccountsUsecase.GetAll()

		assert.NotNil(t, err)
		assert.Equal(t, []payment_accounts.Domain{}, resp)
	})
}

func TestInsert(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockPaymentAccountRepo.On("Insert", mock.Anything).
			Return(paymentAccountsData, nil).Once()

		resp, err := paymentAccountsUsecase.Insert(&paymentAccountsData)

		assert.Nil(t, err)
		assert.Equal(t, paymentAccountsData, resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockPaymentAccountRepo.On("Insert", mock.Anything).
			Return(payment_accounts.Domain{}, assert.AnError).Once()

		resp, err := paymentAccountsUsecase.Insert(&paymentAccountsData)

		assert.NotNil(t, err)
		assert.Equal(t, payment_accounts.Domain{}, resp)
	})
}