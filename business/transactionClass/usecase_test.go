package transactionClass_test

import (
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

	m.Run()
}

func TestInsert(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockTransactionClassRepo.On("Insert", mock.Anything).Return(transactionClassData, nil).Once()
		mockClassRepo.On("UpdateParticipant", mock.AnythingOfType("int")).Return(transactionClassData, nil).Once()

		resp, err := transactionClassUsecase.Insert(&transactionClassInput)

		assert.Nil(t, err)
		assert.Equal(t, transactionClassData, resp)
	})
	// t.Run("Invalid Test | Duplicate Data Error", func(t *testing.T) {
	// 	mockClassRepo.On("Insert", mock.Anything).Return(transactionClass.Domain{}, assert.AnError).Once()
	// 	resp, err := transactionClassUsecase.Insert(&transactionClassInput)

	// 	// assert.NotNil(t, err)
	// 	assert.Equal(t, "Data already exist", err)
	// 	assert.Contains(t, transactionClass.Domain{}, resp)
	// })
}

// func TestGetAll(t *testing.T) {
// 	t.Run("Valid Test | Unspecified Page", func(t *testing.T) {
// 		mockClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
// 			Return([]class.Domain{classData}, int64(1), nil).Once()

// 		expectOffset := 0
// 		expectLimit := 10
// 		expectTotalData := int64(1)
// 		resp, offset, limit, totalData, err := classUsecase.GetAll("Test", 1)

// 		assert.Nil(t, err)
// 		assert.Contains(t, resp, classData)
// 		assert.Equal(t, expectLimit, limit)
// 		assert.Equal(t, expectOffset, offset)
// 		assert.Equal(t, expectTotalData, totalData)
// 	})
// 	t.Run("Valid Test | Specified Page", func(t *testing.T) {
// 		mockClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
// 			Return([]class.Domain{classData}, int64(1), nil).Once()

// 		expectOffset := 10
// 		expectLimit := 10
// 		expectTotalData := int64(1)
// 		resp, offset, limit, totalData, err := classUsecase.GetAll("Test", 2)

// 		assert.Nil(t, err)
// 		assert.Contains(t, resp, classData)
// 		assert.Equal(t, expectLimit, limit)
// 		assert.Equal(t, expectOffset, offset)
// 		assert.Equal(t, expectTotalData, totalData)
// 	})

// 	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
// 		mockClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
// 			Return([]class.Domain{}, int64(0), assert.AnError).Once()

// 		expectOffset := -1
// 		expectLimit := -1
// 		expectTotalData := int64(-1)
// 		resp, offset, limit, totalData, err := classUsecase.GetAll("Test", 1)

// 		assert.NotNil(t, err)
// 		assert.Equal(t, resp, []class.Domain{})
// 		assert.Equal(t, expectLimit, limit)
// 		assert.Equal(t, expectOffset, offset)
// 		assert.Equal(t, expectTotalData, totalData)
// 	})

// }

// func TestUpdateParticipant(t *testing.T) {
// 	t.Run("Valid Test", func(t *testing.T) {
// 		mockClassRepo.On("UpdateParticipant", mock.Anything).Return(classData, nil).Once()

// 		resp, err := classUsecase.UpdateParticipant(1)

// 		assert.Nil(t, err)
// 		assert.Equal(t, "", resp)
// 	})
// 	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
// 		mockClassRepo.On("UpdateParticipant", mock.Anything).Return(class.Domain{}, assert.AnError).Once()
// 		resp, err := classUsecase.UpdateParticipant(1)

// 		assert.NotNil(t, err)
// 		assert.Equal(t, "", resp)
// 	})
// }

// func TestUpdateClassByID(t *testing.T) {
// 	t.Run("Valid Test", func(t *testing.T) {
// 		mockClassRepo.On("UpdateClassByID", mock.Anything, mock.Anything).Return(classData, nil).Once()

// 		resp, err := classUsecase.UpdateClassByID(uint(1), &classInput)

// 		assert.Nil(t, err)
// 		assert.Equal(t, "item edited", resp)
// 	})
// 	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
// 		mockClassRepo.On("UpdateClassByID", mock.Anything, mock.Anything).Return(class.Domain{}, assert.AnError).Once()
// 		resp, err := classUsecase.UpdateClassByID(uint(1), &classInput)

// 		assert.NotNil(t, err)
// 		assert.Equal(t, "", resp)
// 	})
// }
