package members_test

import (
	"gym-membership/business"
	"gym-membership/business/members"
	_memberMock "gym-membership/business/members/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	mockMemberRepo _memberMock.Repository
	memberUsecase  members.Usecase
	memberData     members.Domain
)

func TestMain(m *testing.M) {
	memberUsecase = members.NewMemberUsecase(&mockMemberRepo)
	memberData = members.Domain{
		ID:         1,
		UserID:     1,
		ExpireDate: time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:  time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC),
	}
	m.Run()
}

func TestGetByUserID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMemberRepo.On("GetByUserID", mock.AnythingOfType("uint")).
			Return(memberData, nil).Once()

		resp, err := memberUsecase.GetByUserID(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, memberData, resp)
	})
	t.Run("Invalid Test | Not Found", func(t *testing.T){
		mockMemberRepo.On("GetByUserID", mock.AnythingOfType("uint")).
			Return(members.Domain{}, gorm.ErrRecordNotFound).Once()

		resp, err := memberUsecase.GetByUserID(uint(1))

		assert.NotNil(t, err)
		assert.Equal(t, members.Domain{}, resp)
		assert.Equal(t, business.ErrUserNotFound, err)
	})
	t.Run("Invalid Test | Not Found", func(t *testing.T){
		mockMemberRepo.On("GetByUserID", mock.AnythingOfType("uint")).
			Return(members.Domain{}, assert.AnError).Once()

		resp, err := memberUsecase.GetByUserID(uint(1))

		assert.NotNil(t, err)
		assert.Equal(t, members.Domain{}, resp)
		assert.Equal(t, business.ErrInternalServer, err)
	})
}