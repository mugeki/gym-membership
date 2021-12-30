package trainers_test

import (
	"gym-membership/business/trainers"
	_trainersMock "gym-membership/business/trainers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockTrainersRepo _trainersMock.Repository
	trainersUsecase  trainers.Usecase
	trainersData     trainers.Domain
)

func TestMain(m *testing.M) {
	trainersUsecase = trainers.NewTrainerUsecase(&mockTrainersRepo)
	trainersData = trainers.Domain{
		ID:          1,
		Fullname:    "test string",
		UrlImage:    "test string",
		CreatedAt:   time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
		DeletededAt: time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Valid Test | Unspecified Page", func(t *testing.T) {
		mockTrainersRepo.On("GetAll").
			Return([]trainers.Domain{trainersData}, nil).Once()
		resp, err := trainersUsecase.GetAll()

		assert.Nil(t, err)
		assert.Contains(t, resp, trainersData)
	})

	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockTrainersRepo.On("GetAll").
			Return([]trainers.Domain{}, assert.AnError).Once()

		resp, err := trainersUsecase.GetAll()

		assert.NotNil(t, err)
		assert.Equal(t, resp, []trainers.Domain{})
	})

}
