package users_test

import (
	"gym-membership/app/middleware"
	"gym-membership/business/users"
	_userMock "gym-membership/business/users/mocks"
	"gym-membership/helper/encrypt"
	_encryptMock "gym-membership/helper/encrypt/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockUserRepo	_userMock.Repository
	mockEncrypt		_encryptMock.Helper
	userUsecase		users.Usecase
	userData		users.Domain
	hashedPassword	string
	userUUID		uuid.UUID
)

func TestMain(m *testing.M){
	userUsecase = users.NewUserUsecase(&mockUserRepo, &middleware.ConfigJWT{})
	hashedPassword, _ = encrypt.Hash("testpassword")
	userUUID = uuid.New()
	userData = users.Domain{
		ID 			: 1,
		UUID 		: userUUID,
		Username	: "test123",
		Password	: hashedPassword,
		Email		: "test@gmail.com",
		FullName 	: "Test Name",
		Gender 		: "Male",
		Telephone 	: "88888000102",
		Address 	: "Test Street",
		CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
	}
	m.Run()
}

func TestRegister(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockUserRepo.On("Register", mock.Anything).Return(userData,nil).Once()

		inputData := users.Domain{
			Username	: "test123",
			Password	: "testpassword",
			Email		: "test@gmail.com",
			FullName 	: "Test Name",
			Gender 		: "Male",
			Telephone 	: "88888000102",
			Address 	: "Test Street",
			CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
		}

		resp, err := userUsecase.Register(&inputData)

		assert.Nil(t, err)
		assert.Equal(t, "item created", resp)
	})
	t.Run("Invalid Test | Duplicate Data", func(t *testing.T){
		mockUserRepo.On("Register", mock.Anything).Return(users.Domain{},assert.AnError).Once()

		inputData := users.Domain{
			Username	: "test123",
			Password	: "testpassword",
			Email		: "test@gmail.com",
			FullName 	: "Test Name",
			Gender 		: "Male",
			Telephone 	: "88888000102",
			Address 	: "Test Street",
			CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
		}

		resp, err := userUsecase.Register(&inputData)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestLogin(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockUserRepo.On("GetByUsername", mock.AnythingOfType("string")).
					Return(userData,nil).Once()
		
		username := "test123"
		password := "testpassword"
		
		resp, err := userUsecase.Login(username, password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
	t.Run("Invalid Test | Invalid Username", func(t *testing.T){
		mockUserRepo.On("GetByUsername", mock.AnythingOfType("string")).
					Return(users.Domain{},assert.AnError).Once()
		
		username := "test12"
		password := "testpassword"
		
		resp, err := userUsecase.Login(username, password)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
	t.Run("Invalid Test | Invalid Password", func(t *testing.T){
		mockUserRepo.On("GetByUsername", mock.AnythingOfType("string")).
					Return(userData,nil).Once()
		mockEncrypt.On("ValidateHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(false).Once()
		
		username := "test123"
		password := "testpasswor"
		
		resp, err := userUsecase.Login(username, password)

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}