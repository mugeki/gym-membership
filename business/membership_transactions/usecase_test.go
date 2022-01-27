package membership_transactions_test

import (
	"gym-membership/business"
	_memberMock "gym-membership/business/members/mocks"
	"gym-membership/business/membership_products"
	_membershipProductMock "gym-membership/business/membership_products/mocks"
	"gym-membership/business/membership_transactions"
	_membershipTransactionMock "gym-membership/business/membership_transactions/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	mockMembershipTransactionRepo _membershipTransactionMock.Repository
	mockMembershipProductRepo _membershipProductMock.Repository
	mockMemberRepo	_memberMock.Repository
	membershipTransactionUsecase  membership_transactions.Usecase
	membershipTransactionData     membership_transactions.Domain
	membershipTransactionInput    membership_transactions.Domain
	membershipTransactionUpdate  membership_transactions.Domain
	productData                membership_products.Domain
)

func TestMain(m *testing.M) {
	membershipTransactionUsecase = membership_transactions.NewMembershipTransactionUsecase(&mockMembershipTransactionRepo, &mockMembershipProductRepo, &mockMemberRepo)
	membershipTransactionData = membership_transactions.Domain{
		ID:      1,
		UserID:  1,
		AdminID: 1,
		Status:  "waiting for payment",
		Nominal: 100000,
		MembershipProductID: 1,
		CreatedAt:    time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	}
	membershipTransactionInput = membership_transactions.Domain{
		UserID:  1,
		AdminID: 1,
		Status:  "waiting for payment",
		Nominal: 100000,
		MembershipProductID: 1,
	}
	membershipTransactionUpdate = membership_transactions.Domain{
		ID:      1,
		UserID:  1,
		AdminID: 1,
		Status:  "failed",
		Nominal: 100000,
		MembershipProductID: 1,
		CreatedAt:    time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	}
	productData = membership_products.Domain{
		ID         : 1,
		Name       : "test product",
		UrlImage   : "testurl",
		Price      : 200000,
		PeriodTime : 30,
	}

	m.Run()
}

func TestInsert(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockMembershipTransactionRepo.On("Insert", mock.Anything).Return(membershipTransactionData, nil).Once()

		resp, err := membershipTransactionUsecase.Insert(&membershipTransactionInput)

		assert.Nil(t, err)
		assert.Equal(t, membershipTransactionData, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockMembershipTransactionRepo.On("Insert", mock.Anything).Return(membership_transactions.Domain{}, assert.AnError).Once()
		
		resp, err := membershipTransactionUsecase.Insert(&membershipTransactionInput)

		assert.NotNil(t, err)
		assert.Equal(t, membership_transactions.Domain{}, resp)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Valid Test | Unspecified Page", func(t *testing.T) {
		mockMembershipTransactionRepo.On("GetAll", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]membership_transactions.Domain{membershipTransactionData}, int64(1), nil).Once()

		expectOffset := 0
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := membershipTransactionUsecase.GetAll(time.Time{}, "Test", uint(1), 1)

		assert.Nil(t, err)
		assert.Contains(t, resp, membershipTransactionData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Valid Test | Specified Page", func(t *testing.T) {
		mockMembershipTransactionRepo.On("GetAll", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]membership_transactions.Domain{membershipTransactionData}, int64(1), nil).Once()

		expectOffset := 10
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := membershipTransactionUsecase.GetAll(time.Time{}, "Test", uint(1), 2)

		assert.Nil(t, err)
		assert.Contains(t, resp, membershipTransactionData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})

	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockMembershipTransactionRepo.On("GetAll", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]membership_transactions.Domain{}, int64(0), assert.AnError).Once()

		expectOffset := -1
		expectLimit := -1
		expectTotalData := int64(-1)
		resp, offset, limit, totalData, err := membershipTransactionUsecase.GetAll(time.Time{}, "Test", uint(1), 2)

		assert.NotNil(t, err)
		assert.Equal(t, resp, []membership_transactions.Domain{})
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})

}

func TestUpdateStatus(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockMembershipTransactionRepo.On("UpdateStatus", mock.AnythingOfType("uint"),  mock.AnythingOfType("uint"), mock.AnythingOfType("string")).
			Return(membershipTransactionData, nil).Once()
		mockMembershipProductRepo.On("GetByID", mock.AnythingOfType("uint")).Return(productData, nil).Once()

		err := membershipTransactionUsecase.UpdateStatus(uint(1), uint(1), "waiting-for-confirmation")

		assert.Nil(t, err)
	})
	t.Run("Valid Test (Transaction Accepted)", func(t *testing.T) {
		mockMembershipTransactionRepo.On("UpdateStatus", mock.AnythingOfType("uint"),  mock.AnythingOfType("uint"), mock.AnythingOfType("string")).
			Return(membershipTransactionData, nil).Once()
		mockMembershipProductRepo.On("GetByID", mock.AnythingOfType("uint")).Return(productData, nil).Once()
		mockMemberRepo.On("Insert", mock.Anything).Return(nil).Once()

		err := membershipTransactionUsecase.UpdateStatus(uint(1), uint(1), "accepted")

		assert.Nil(t, err)
	})
	t.Run("Invalid Test | Internal Server Error (UpdateStatus)", func(t *testing.T) {
		mockMembershipTransactionRepo.On("UpdateStatus", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("string")).
			Return(membership_transactions.Domain{}, assert.AnError).Once()

		err := membershipTransactionUsecase.UpdateStatus(uint(1), uint(1), "accepted")

		assert.NotNil(t, err)
	})
	t.Run("Invalid Test | Internal Server Error (GetByID)", func(t *testing.T) {
		mockMembershipTransactionRepo.On("UpdateStatus", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("string")).
			Return(membershipTransactionData, nil).Once()
		mockMembershipProductRepo.On("GetByID", mock.AnythingOfType("uint")).Return(membership_products.Domain{}, assert.AnError).Once()

		err := membershipTransactionUsecase.UpdateStatus(uint(1), uint(1), "accepted")

		assert.NotNil(t, err)
	})
	t.Run("Invalid Test | Product Not Found", func(t *testing.T) {
		mockMembershipTransactionRepo.On("UpdateStatus", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("string")).
			Return(membershipTransactionData, nil).Once()
		mockMembershipProductRepo.On("GetByID", mock.AnythingOfType("uint")).Return(membership_products.Domain{}, gorm.ErrRecordNotFound).Once()

		err := membershipTransactionUsecase.UpdateStatus(uint(1), uint(1), "accepted")

		assert.NotNil(t, err)
		assert.Equal(t, business.ErrProductNotFound, err)
	})
	t.Run("Invalid Test | Internal Server Error (Insert)", func(t *testing.T) {
		mockMembershipTransactionRepo.On("UpdateStatus", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("string")).
			Return(membershipTransactionData, nil).Once()
		mockMembershipProductRepo.On("GetByID", mock.AnythingOfType("uint")).Return(productData, nil).Once()
		mockMemberRepo.On("Insert", mock.Anything).Return(assert.AnError).Once()

		err := membershipTransactionUsecase.UpdateStatus(uint(1), uint(1), "accepted")

		assert.NotNil(t, err)
	})
}

func TestGetAllByUser(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMembershipTransactionRepo.On("GetAllByUser", mock.AnythingOfType("uint")).
			Return([]membership_transactions.Domain{membershipTransactionData} , nil).Once()

		resp, err := membershipTransactionUsecase.GetAllByUser(uint(1))

		assert.Nil(t, err)
		assert.Contains(t, resp, membershipTransactionData)
		assert.Len(t, resp, 1)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockMembershipTransactionRepo.On("GetAllByUser", mock.AnythingOfType("uint")).
			Return([]membership_transactions.Domain{} , assert.AnError).Once()

		resp, err := membershipTransactionUsecase.GetAllByUser(uint(1))

		assert.NotNil(t, err)
		assert.Equal(t, []membership_transactions.Domain{}, resp)
	})
}

func TestGetByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMembershipTransactionRepo.On("GetByID", mock.AnythingOfType("uint")).
			Return(membershipTransactionData , nil).Once()

		resp, err := membershipTransactionUsecase.GetByID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, membershipTransactionData, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockMembershipTransactionRepo.On("GetByID", mock.AnythingOfType("uint")).
			Return(membership_transactions.Domain{} , assert.AnError).Once()

		resp, err := membershipTransactionUsecase.GetByID(uint(1))

		assert.NotNil(t, err)
		assert.Equal(t, membership_transactions.Domain{}, resp)
	})
}

func TestUpdateStatusToFailed(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMembershipTransactionRepo.On("UpdateStatusToFailed", mock.AnythingOfType("uint"), 
			mock.AnythingOfType("string")).Return(membershipTransactionInput, nil).Once()
		
		resp, err := membershipTransactionUsecase.UpdateStatusToFailed(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, membershipTransactionInput, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockMembershipTransactionRepo.On("UpdateStatusToFailed", mock.AnythingOfType("uint"), 
			mock.AnythingOfType("string")).Return(membership_transactions.Domain{}, assert.AnError).Once()
		
		resp, err := membershipTransactionUsecase.UpdateStatusToFailed(uint(1))

		assert.NotNil(t, err)
		assert.Equal(t, business.ErrInternalServer, err)
		assert.Equal(t, membership_transactions.Domain{}, resp)
	})
}