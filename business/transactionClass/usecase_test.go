package transactionClass_test

import (
	"gym-membership/business/class"
	_classMock "gym-membership/business/class/mocks"
	"gym-membership/business/transactionClass"
	_transactionClassMock "gym-membership/business/transactionClass/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockTransactionClassRepo _transactionClassMock.Repository
	mockClassRepo            _classMock.Repository
	transactionClassUsecase  transactionClass.Usecase
	transactionClassData     transactionClass.Domain
	transactionClassInput    transactionClass.Domain
	classData                class.Domain
)

func TestMain(m *testing.M) {
	transactionClassUsecase = transactionClass.NewTransactionClassUsecase(&mockTransactionClassRepo, &mockClassRepo)
	transactionClassData = transactionClass.Domain{
		ID:      1,
		UserID:  1,
		AdminID: 1,
		Status:  "waiting for payment",
		Nominal: 100000,
		ClassID: 2,
		Date:    time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	}
	transactionClassInput = transactionClass.Domain{
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
		mockTransactionClassRepo.On("Insert", mock.Anything).Return(transactionClassData, nil).Once()
		mockClassRepo.On("UpdateParticipant", mock.AnythingOfType("int")).Return(classData, nil).Once()

		resp, err := transactionClassUsecase.Insert(&transactionClassInput)

		assert.Nil(t, err)
		assert.Equal(t, transactionClassData, resp)
	})
	t.Run("Invalid Test | Duplicate Data Error", func(t *testing.T) {
		mockTransactionClassRepo.On("Insert", mock.Anything).Return(transactionClass.Domain{}, assert.AnError).Once()
		resp, err := transactionClassUsecase.Insert(&transactionClassInput)

		assert.NotNil(t, err)
		assert.Equal(t, transactionClass.Domain{}, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockTransactionClassRepo.On("Insert", mock.Anything).Return(transactionClassData, nil).Once()
		mockClassRepo.On("UpdateParticipant", mock.AnythingOfType("int")).Return(class.Domain{}, assert.AnError).Once()
		resp, err := transactionClassUsecase.Insert(&transactionClassInput)

		assert.NotNil(t, err)
		assert.Equal(t, transactionClass.Domain{}, resp)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Valid Test | Unspecified Page", func(t *testing.T) {
		mockTransactionClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]transactionClass.Domain{transactionClassData}, int64(1), nil).Once()

		expectOffset := 0
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := transactionClassUsecase.GetAll("Test", uint(1), 1)

		assert.Nil(t, err)
		assert.Contains(t, resp, transactionClassData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Valid Test | Specified Page", func(t *testing.T) {
		mockTransactionClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]transactionClass.Domain{transactionClassData}, int64(1), nil).Once()

		expectOffset := 10
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := transactionClassUsecase.GetAll("Test", uint(1), 2)

		assert.Nil(t, err)
		assert.Contains(t, resp, transactionClassData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})

	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockTransactionClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]transactionClass.Domain{}, int64(0), assert.AnError).Once()

		expectOffset := -1
		expectLimit := -1
		expectTotalData := int64(-1)
		resp, offset, limit, totalData, err := transactionClassUsecase.GetAll("Test", uint(1), 2)

		assert.NotNil(t, err)
		assert.Equal(t, resp, []transactionClass.Domain{})
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})

}

func TestUpdateStatus(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockTransactionClassRepo.On("UpdateStatus", mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(transactionClassData, nil).Once()

		resp, err := transactionClassUsecase.UpdateStatus(uint(1), "accepted")

		assert.Nil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockTransactionClassRepo.On("UpdateStatus", mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(transactionClass.Domain{}, assert.AnError).Once()
		resp, err := transactionClassUsecase.UpdateStatus(uint(1), "accepted")

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}
