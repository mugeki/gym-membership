package admins_test

import (
	"gym-membership/app/middleware"
	"gym-membership/business/admins"
	_adminMock "gym-membership/business/admins/mocks"
	"gym-membership/helper/encrypt"
	_encryptMock "gym-membership/helper/encrypt/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockAdminRepo		_adminMock.Repository
	mockEncrypt			_encryptMock.Helper
	adminUsecase		admins.Usecase
	adminData			admins.Domain
	hashedPassword		string
)

func TestMain(m *testing.M){
	adminUsecase = admins.NewAdminUsecase(&mockAdminRepo, &middleware.ConfigJWT{})
	hashedPassword, _ = encrypt.Hash("testpassword")
	adminData = admins.Domain{
		ID 			: 1,
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
		mockAdminRepo.On("Register", mock.Anything).Return(adminData,nil).Once()

		inputData := admins.Domain{
			Username	: "test123",
			Password	: "testpassword",
			Email		: "test@gmail.com",
			FullName 	: "Test Name",
			Gender 		: "Male",
			Telephone 	: "88888000102",
			Address 	: "Test Street",
			CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
		}

		resp, err := adminUsecase.Register(&inputData)

		assert.Nil(t, err)
		assert.Equal(t, adminData, resp)
	})
	t.Run("Invalid Test | Duplicate Data", func(t *testing.T){
		mockAdminRepo.On("Register", mock.Anything).Return(admins.Domain{},assert.AnError).Once()

		inputData := admins.Domain{
			Username	: "test123",
			Password	: "testpassword",
			Email		: "test@gmail.com",
			FullName 	: "Test Name",
			Gender 		: "Male",
			Telephone 	: "88888000102",
			Address 	: "Test Street",
			CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
		}

		resp, err := adminUsecase.Register(&inputData)

		assert.NotNil(t, err)
		assert.Equal(t, admins.Domain{}, resp)
	})
}

func TestLogin(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockAdminRepo.On("GetByUsername", mock.AnythingOfType("string")).
			Return(adminData,nil).Once()

		adminname := "test123"
		password := "testpassword"
		
		resp, err := adminUsecase.Login(adminname, password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
	t.Run("Valid Test (Is Member)", func(t *testing.T){
		mockAdminRepo.On("GetByUsername", mock.AnythingOfType("string")).
			Return(adminData,nil).Once()

		adminname := "test123"
		password := "testpassword"
		
		resp, err := adminUsecase.Login(adminname, password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
	t.Run("Invalid Test | Invalid Username", func(t *testing.T){
		mockAdminRepo.On("GetByUsername", mock.AnythingOfType("string")).
					Return(admins.Domain{},assert.AnError).Once()
		
		adminname := "test12"
		password := "testpassword"
		
		resp, err := adminUsecase.Login(adminname, password)

		assert.NotNil(t, err)
		assert.Equal(t, admins.Domain{}, resp)
	})
	t.Run("Invalid Test | Invalid Password", func(t *testing.T){
		mockAdminRepo.On("GetByUsername", mock.AnythingOfType("string")).
			Return(adminData,nil).Once()
		mockEncrypt.On("ValidateHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(false).Once()
		
		adminname := "test123"
		password := "testpasswor"
		
		resp, err := adminUsecase.Login(adminname, password)

		assert.NotNil(t, err)
		assert.Equal(t, admins.Domain{}, resp)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockAdminRepo.On("Update", mock.AnythingOfType("uint"), mock.Anything).
			Return(adminData,nil).Once()

		inputData := admins.Domain{
			Username	: "test123",
			Password	: "testpassword",
			Email		: "test@gmail.com",
			FullName 	: "Test Name",
			Gender 		: "Male",
			Telephone 	: "88888000102",
			Address 	: "Test Street",
			CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
		}

		res, err := adminUsecase.Update(uint(1), &inputData)

		assert.Nil(t, err)
		assert.Equal(t, adminData, res)
	})
	t.Run("Valid Test | No Password", func(t *testing.T){
		mockAdminRepo.On("GetByID", mock.AnythingOfType("uint")).
			Return(adminData, nil).Once()
		mockAdminRepo.On("Update", mock.AnythingOfType("uint"), mock.Anything).
			Return(admins.Domain{},assert.AnError).Once()

		inputData := admins.Domain{
			Username	: "test123",
			Email		: "test@gmail.com",
			FullName 	: "Test Name",
			Gender 		: "Male",
			Telephone 	: "88888000102",
			Address 	: "Test Street",
			CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
		}

		res, err := adminUsecase.Update(uint(1), &inputData)

		assert.NotNil(t, err)
		assert.Equal(t, admins.Domain{}, res)
	})
	t.Run("Valid Test | Internal Server Error 1", func(t *testing.T){
		mockAdminRepo.On("GetByID", mock.AnythingOfType("uint")).
			Return(admins.Domain{}, assert.AnError).Once()
		mockAdminRepo.On("Update", mock.AnythingOfType("uint"), mock.Anything).
			Return(admins.Domain{},assert.AnError).Once()

		inputData := admins.Domain{
			Username	: "test123",
			Email		: "test@gmail.com",
			FullName 	: "Test Name",
			Gender 		: "Male",
			Telephone 	: "88888000102",
			Address 	: "Test Street",
			CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
		}

		res, err := adminUsecase.Update(uint(1), &inputData)

		assert.NotNil(t, err)
		assert.Equal(t, admins.Domain{}, res)
	})
	t.Run("Invalid Test | Internal Server Error 2", func(t *testing.T){
		mockAdminRepo.On("Update", mock.AnythingOfType("uint"), mock.Anything).
			Return(admins.Domain{},assert.AnError).Once()

		inputData := admins.Domain{
			Username	: "test123",
			Password	: "testpassword",
			Email		: "test@gmail.com",
			FullName 	: "Test Name",
			Gender 		: "Male",
			Telephone 	: "88888000102",
			Address 	: "Test Street",
			CreatedAt 	: time.Date(2021,12,1,0,0,0,0,time.UTC),
		}

		res, err := adminUsecase.Update(uint(1), &inputData)

		assert.NotNil(t, err)
		assert.Equal(t, admins.Domain{}, res)
	})
}

func TestGetAll(t *testing.T){
	t.Run("Valid Test | Unspecified Page", func(t *testing.T){
		mockAdminRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]admins.Domain{adminData}, int64(1), nil).Once()

		expectOffset	:= 0
		expectLimit		:= 10
		expectTotalData	:= int64(1)
		resp, offset, limit, totalData, err := adminUsecase.GetAll("Test",1)
		
		assert.Nil(t, err)
		assert.Contains(t, resp, adminData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Valid Test | Specified Page", func(t *testing.T){
		mockAdminRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]admins.Domain{adminData}, int64(1), nil).Once()

		expectOffset	:= 10
		expectLimit		:= 10
		expectTotalData	:= int64(1)
		resp, offset, limit, totalData, err := adminUsecase.GetAll("Test",2)
		
		assert.Nil(t, err)
		assert.Contains(t, resp, adminData)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockAdminRepo.On("GetAll", mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return([]admins.Domain{}, int64(0), assert.AnError).Once()

		expectOffset	:= -1
		expectLimit		:= -1
		expectTotalData	:= int64(-1)
		resp, offset, limit, totalData, err := adminUsecase.GetAll("Test",1)

		assert.NotNil(t, err)
		assert.Equal(t, []admins.Domain{}, resp)
		assert.Equal(t, expectLimit, limit)
		assert.Equal(t, expectOffset, offset)
		assert.Equal(t, expectTotalData, totalData)
	})
}