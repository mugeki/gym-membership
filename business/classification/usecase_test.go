package classification_test

import (
	"gym-membership/business/classification"
	_classificationMock "gym-membership/business/classification/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockClassificationRepo _classificationMock.Repository
	classificationUsecase  classification.Usecase
	classificationData     classification.Domain
)

func TestMain(m *testing.M) {
	classificationUsecase = classification.NewClassificationUsecase(&mockClassificationRepo)
	classificationData = classification.Domain{
		ID:         1,
		Name:		"Test Classification",
	}
	m.Run()
}

func TestGetAll(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockClassificationRepo.On("GetAll").
			Return([]classification.Domain{classificationData}, nil).Once()

		resp, err := classificationUsecase.GetAll()

		assert.Nil(t, err)
		assert.Contains(t, resp, classificationData)
		assert.Len(t, resp, 1)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockClassificationRepo.On("GetAll").
			Return([]classification.Domain{}, assert.AnError).Once()

		resp, err := classificationUsecase.GetAll()

		assert.NotNil(t, err)
		assert.Equal(t, []classification.Domain{}, resp)
	})
}

func TestInsert(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockClassificationRepo.On("Insert", mock.Anything).
			Return(classificationData, nil).Once()

		resp, err := classificationUsecase.Insert(&classificationData)

		assert.Nil(t, err)
		assert.Equal(t, classificationData, resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockClassificationRepo.On("Insert", mock.Anything).
			Return(classification.Domain{}, assert.AnError).Once()

		resp, err := classificationUsecase.Insert(&classificationData)

		assert.NotNil(t, err)
		assert.Equal(t, classification.Domain{}, resp)
	})
}