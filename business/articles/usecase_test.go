package articles_test

import (
	"gym-membership/business/articles"
	_articleMock "gym-membership/business/articles/mocks"
	_classificationMock "gym-membership/business/classification/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockArticleRepo        _articleMock.Repository
	mockClassificationRepo _classificationMock.Repository
	articleUsecase         articles.Usecase
	articleData            articles.Domain
	articleInput           articles.Domain
)

func TestMain(m *testing.M) {
	articleUsecase = articles.NewArticleUsecase(&mockArticleRepo, &mockClassificationRepo)
	articleData = articles.Domain{
		ID:               1,
		Title:            "Test Video",
		ClassificationID: 1,
		AdminID:          1,
		MemberOnly:       false,
		UrlImage:         "https://www.image.com/watch?v=dQw4w9WgXcQ",
		CreatedAt:        time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
		Text:             "this is a long text for testing",
	}
	articleInput = articles.Domain{
		Title:            "Test Video",
		ClassificationID: 1,
		MemberOnly:       false,
		UrlImage:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Valid Test | Unspecified Page", func(t *testing.T) {
		mockArticleRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]articles.Domain{articleData}, int64(1), nil).Once()

		expectOffset := 0
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := articleUsecase.GetAll("Test", 1)

		assert.Nil(t, err)
		assert.Contains(t, resp, articleData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Valid Test | Specified Page", func(t *testing.T) {
		mockArticleRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]articles.Domain{articleData}, int64(1), nil).Once()

		expectOffset := 10
		expectLimit := 10
		expectTotalData := int64(1)
		resp, offset, limit, totalData, err := articleUsecase.GetAll("Test", 2)

		assert.Nil(t, err)
		assert.Contains(t, resp, articleData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockArticleRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]articles.Domain{}, int64(0), assert.AnError).Once()

		expectOffset := -1
		expectLimit := -1
		expectTotalData := int64(-1)
		resp, offset, limit, totalData, err := articleUsecase.GetAll("Test", 1)

		assert.NotNil(t, err)
		assert.Equal(t, []articles.Domain{}, resp)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockArticleRepo.On("GetClassificationID", mock.AnythingOfType("string")).Return(1, nil).Once()
		mockArticleRepo.On("Insert", mock.Anything).Return(articleData, nil).Once()

		resp, err := articleUsecase.Insert(&articleInput)

		assert.Nil(t, err)
		assert.Equal(t, articleData, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockArticleRepo.On("GetClassificationID", mock.AnythingOfType("string")).Return(1, nil).Once()
		mockArticleRepo.On("Insert", mock.Anything).Return(articles.Domain{}, assert.AnError).Once()

		resp, err := articleUsecase.Insert(&articleInput)

		assert.NotNil(t, err)
		assert.Equal(t, articles.Domain{}, resp)
	})
}

func TestUpdateByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockArticleRepo.On("GetClassificationID", mock.AnythingOfType("string")).Return(1, nil).Once()
		mockArticleRepo.On("UpdateByID", mock.Anything, mock.Anything).Return(articleData, nil).Once()

		resp, err := articleUsecase.UpdateArticleByID(1, &articleInput)

		assert.Nil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		mockArticleRepo.On("GetClassificationID", mock.AnythingOfType("string")).Return(1, nil).Once()
		mockArticleRepo.On("UpdateByID", mock.Anything, mock.Anything).Return(articles.Domain{}, assert.AnError).Once()

		resp, err := articleUsecase.UpdateArticleByID(1, &articleInput)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}
