package class_test

import (
	"gym-membership/app/middleware"
	"gym-membership/business/class"
	_classMock "gym-membership/business/class/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockClassRepo _classMock.Repository
	classUsecase  class.Usecase
	classData     class.Domain
	classInput    class.Domain
)

func TestMain(m *testing.M) {
	classUsecase = class.NewClassUsecase(&mockClassRepo, &middleware.ConfigJWT{})
	classData = class.Domain{
		ID:              1,
		Name:            "Test class",
		Price:           100000,
		Kuota:           20,
		Participant:     15,
		TrainerId:       1,
		TrainerName:     "jono",
		TrainerImage:    "www.image.com",
		Description:     "description test",
		AvailableStatus: false,
		IsOnline:        true,
		Date:            "2021-08-10T15:00:00 ; 2021-08-10T15:00:00 ",
		Location:        "google-meet/join.com",

		CreatedAt: time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	}
	classInput = class.Domain{
		Name:        "Test class",
		Price:       100000,
		Kuota:       20,
		Participant: 15,
		TrainerId:   1,
	}

	m.Run()
}

func TestInsert(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockClassRepo.On("Insert", mock.Anything).Return(classData, nil).Once()

		resp, err := classUsecase.Insert(&classInput)

		assert.Nil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockClassRepo.On("Insert", mock.Anything).Return(class.Domain{}, assert.AnError).Once()
		resp, err := classUsecase.Insert(&classInput)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Valid Test | Unspecified Page", func(t *testing.T) {
		mockClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]class.Domain{classData}, int64(1), nil).Once()

		expectOffset := 0
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := classUsecase.GetAll("Test", 1)

		assert.Nil(t, err)
		assert.Contains(t, resp, classData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Valid Test | Specified Page", func(t *testing.T) {
		mockClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]class.Domain{classData}, int64(1), nil).Once()

		expectOffset := 10
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := classUsecase.GetAll("Test", 2)

		assert.Nil(t, err)
		assert.Contains(t, resp, classData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})

	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockClassRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]class.Domain{}, int64(0), assert.AnError).Once()

		expectOffset := -1
		expectLimit := -1
		expectTotalData := int64(-1)
		resp, offset, limit, totalData, err := classUsecase.GetAll("Test", 1)

		assert.NotNil(t, err)
		assert.Equal(t, resp, []class.Domain{})
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})

}

func TestUpdateParticipant(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockClassRepo.On("UpdateParticipant", mock.Anything).Return(classData, nil).Once()

		resp, err := classUsecase.UpdateParticipant(1)

		assert.Nil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockClassRepo.On("UpdateParticipant", mock.Anything).Return(class.Domain{}, assert.AnError).Once()
		resp, err := classUsecase.UpdateParticipant(1)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestUpdateClassByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockClassRepo.On("UpdateClassByID", mock.Anything, mock.Anything).Return(classData, nil).Once()

		resp, err := classUsecase.UpdateClassByID(uint(1), &classInput)

		assert.Nil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockClassRepo.On("UpdateClassByID", mock.Anything, mock.Anything).Return(class.Domain{}, assert.AnError).Once()
		resp, err := classUsecase.UpdateClassByID(uint(1), &classInput)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}
