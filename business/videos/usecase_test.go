package videos_test

import (
	"gym-membership/business"
	"gym-membership/business/videos"
	_videoMock "gym-membership/business/videos/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockVideoRepo	_videoMock.Repository
	videoUsecase	videos.Usecase
	videoData		videos.Domain
	videoInput		videos.Domain
)

func TestMain(m *testing.M){
	videoUsecase = videos.NewVideoUsecase(&mockVideoRepo)
	videoData = videos.Domain{
		ID                 : 1,
		Title              : "Test Video",
		ClassificationID   : 1,
		ClassificationName : "test name",
		AdminID            : 1,
		MemberOnly         : false,
		Url                : "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		CreatedAt          : time.Date(2021,12,1,0,0,0,0,time.UTC),
	}
	videoInput = videos.Domain{
		Title              : "Test Video",
		ClassificationName : "test name",
		MemberOnly         : false,
		Url                : "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	}
	
	m.Run()
}

func TestGetAll(t *testing.T){
	t.Run("Valid Test | Unspecified Page", func(t *testing.T){
		mockVideoRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
						Return([]videos.Domain{videoData}, int64(1), nil).Once()

		expectOffset	:= 0
		expectLimit		:= 10
		expectTotalData	:= int64(1)
		resp, offset, limit, totalData, err := videoUsecase.GetAll("Test",1)
		
		assert.Nil(t, err)
		assert.Contains(t, resp, videoData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Valid Test | Specified Page", func(t *testing.T){
		mockVideoRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
						Return([]videos.Domain{videoData}, int64(1), nil).Once()

		expectOffset	:= 10
		expectLimit		:= 10
		expectTotalData	:= int64(1)
		resp, offset, limit, totalData, err := videoUsecase.GetAll("Test",2)
		
		assert.Nil(t, err)
		assert.Contains(t, resp, videoData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockVideoRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
						Return([]videos.Domain{}, int64(0), assert.AnError).Once()

		expectOffset	:= -1
		expectLimit		:= -1
		expectTotalData	:= int64(-1)
		resp, offset, limit, totalData, err := videoUsecase.GetAll("Test",1)

		assert.NotNil(t, err)
		assert.Equal(t, []videos.Domain{}, resp)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
}

func TestInsert(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockVideoRepo.On("Insert", mock.Anything).Return(videoData, nil).Once()

		resp, err := videoUsecase.Insert(&videoInput)

		assert.Nil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockVideoRepo.On("Insert", mock.Anything).Return(videos.Domain{}, assert.AnError).Once()
		
		resp, err := videoUsecase.Insert(&videoInput)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestUpdateByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockVideoRepo.On("UpdateByID", mock.Anything, mock.Anything).Return(videoData, nil).Once()

		resp, err := videoUsecase.UpdateByID(1, &videoInput)

		assert.Nil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockVideoRepo.On("UpdateByID", mock.Anything, mock.Anything).Return(videos.Domain{}, assert.AnError).Once()

		resp, err := videoUsecase.UpdateByID(1, &videoInput)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestDeleteByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockVideoRepo.On("DeleteByID", mock.AnythingOfType("uint")).Return(nil).Once()

		err := videoUsecase.DeleteByID(1)

		assert.Nil(t, err)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockVideoRepo.On("DeleteByID", mock.AnythingOfType("uint")).Return(assert.AnError).Once()

		err := videoUsecase.DeleteByID(1)

		assert.NotNil(t, err)
		assert.Equal(t, business.ErrInternalServer, err)
	})
}