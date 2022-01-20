package class_transactions_test

import (
	"gym-membership/business/class"
	_classMock "gym-membership/business/class/mocks"
	"gym-membership/business/class_transactions"
	_classTransactionMock "gym-membership/business/class_transactions/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockClassTransactionRepo _classTransactionMock.Repository
	mockClassRepo            _classMock.Repository
	classTransactionUsecase  class_transactions.Usecase
	classTransactionData     class_transactions.Domain
	classTransactionInput    class_transactions.Domain
	classData                class.Domain
)

func TestMain(m *testing.M) {
	classTransactionUsecase = class_transactions.NewClassTransactionUsecase(&mockClassTransactionRepo, &mockClassRepo)
	classTransactionData = class_transactions.Domain{
		ID:        1,
		UserID:    1,
		AdminID:   1,
		Status:    "waiting for payment",
		Nominal:   100000,
		ClassID:   2,
		CreatedAt: time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	}
	classTransactionInput = class_transactions.Domain{
		UserID:  1,
		AdminID: 1,
		Status:  "waiting for payment",
		Nominal: 100000,
		ClassID: 2,
	}

	classData = class.Domain{
		Name:            "test",
		UrlImage:        "testurl",
		Price:           20000,
		Kuota:           10,
		Participant:     3,
		TrainerId:       1,
		TrainerName:     "name",
		TrainerImage:    "fgjhbhddbfhdjs",
		Description:     "description",
		AvailableStatus: true,
		IsOnline:        true,
		Date:            "10-3-4",
		Location:        "jl dr jalan jalna tes",
	}

	m.Run()
}

func TestInsert(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockClassTransactionRepo.On("Insert", mock.Anything).Return(classTransactionData, nil).Once()
		mockClassRepo.On("UpdateParticipant", mock.AnythingOfType("uint")).Return(classData, nil).Once()

		resp, err := classTransactionUsecase.Insert(&classTransactionInput)

		assert.Nil(t, err)
		assert.Equal(t, classTransactionData, resp)
	})
	t.Run("Invalid Test | Duplicate Data Error", func(t *testing.T) {
		mockClassTransactionRepo.On("Insert", mock.Anything).Return(class_transactions.Domain{}, assert.AnError).Once()
		resp, err := classTransactionUsecase.Insert(&classTransactionInput)

		assert.NotNil(t, err)
		assert.Equal(t, class_transactions.Domain{}, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockClassTransactionRepo.On("Insert", mock.Anything).Return(classTransactionData, nil).Once()
		mockClassRepo.On("UpdateParticipant", mock.AnythingOfType("uint")).Return(class.Domain{}, assert.AnError).Once()
		resp, err := classTransactionUsecase.Insert(&classTransactionInput)

		assert.NotNil(t, err)
		assert.Equal(t, class_transactions.Domain{}, resp)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Valid Test | Unspecified Page", func(t *testing.T) {
		mockClassTransactionRepo.On("GetAll", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]class_transactions.Domain{classTransactionData}, int64(1), nil).Once()

		expectOffset := 0
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := classTransactionUsecase.GetAll(time.Time{}, "Test", uint(1), 1)

		assert.Nil(t, err)
		assert.Contains(t, resp, classTransactionData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Valid Test | Specified Page", func(t *testing.T) {
		mockClassTransactionRepo.On("GetAll", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]class_transactions.Domain{classTransactionData}, int64(1), nil).Once()

		expectOffset := 10
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := classTransactionUsecase.GetAll(time.Time{}, "Test", uint(1), 2)

		assert.Nil(t, err)
		assert.Contains(t, resp, classTransactionData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})

	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockClassTransactionRepo.On("GetAll", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]class_transactions.Domain{}, int64(0), assert.AnError).Once()

		expectOffset := -1
		expectLimit := -1
		expectTotalData := int64(-1)
		resp, offset, limit, totalData, err := classTransactionUsecase.GetAll(time.Time{}, "Test", uint(1), 2)

		assert.NotNil(t, err)
		assert.Equal(t, resp, []class_transactions.Domain{})
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})

}

func TestUpdateStatus(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockClassTransactionRepo.On("UpdateStatus", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(classTransactionData, nil).Once()

		resp, err := classTransactionUsecase.UpdateStatus(uint(1), uint(1), "accepted")

		assert.Nil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockClassTransactionRepo.On("UpdateStatus", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(class_transactions.Domain{}, assert.AnError).Once()
		resp, err := classTransactionUsecase.UpdateStatus(uint(1), uint(1), "accepted")

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestGetActiveClass(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockClassTransactionRepo.On("GetActiveClass", mock.AnythingOfType("uint")).Return([]class.Domain{classData}, nil).Once()
		resp, err := classTransactionUsecase.GetActiveClass(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, []class.Domain{classData}, resp)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockClassTransactionRepo.On("GetActiveClass", mock.AnythingOfType("uint")).Return([]class.Domain{}, assert.AnError).Once()
		resp, err := classTransactionUsecase.GetActiveClass(uint(1))

		assert.NotNil(t, err)
		assert.Equal(t, []class.Domain{}, resp)
	})
}
