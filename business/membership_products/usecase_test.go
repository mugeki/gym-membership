package membership_products_test

import (
	"gym-membership/business"
	"gym-membership/business/membership_products"
	_productMock "gym-membership/business/membership_products/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	mockMembershipProductRepo _productMock.Repository
	productUsecase membership_products.Usecase
	productData membership_products.Domain
)


func TestMain(m *testing.M){
	productUsecase = membership_products.NewMembershipProductsUsecase(&mockMembershipProductRepo)
	productData = membership_products.Domain{
		ID         : 1,
		Name       : "Test Product",
		UrlImage   : "www.image.com",
		Price      : 200000,
		PeriodTime : 30,
	}
	m.Run()
}

func TestInsert(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMembershipProductRepo.On("Insert", mock.Anything).Return(productData, nil).Once()
		
		resp, err := productUsecase.Insert(&productData)

		assert.Nil(t, err)
		assert.Equal(t, productData, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockMembershipProductRepo.On("Insert", mock.Anything).Return(membership_products.Domain{}, assert.AnError).Once()
		
		resp, err := productUsecase.Insert(&productData)

		assert.NotNil(t, err)
		assert.Equal(t, membership_products.Domain{}, resp)
	})
}

func TestGetAll(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMembershipProductRepo.On("GetAll").
			Return([]membership_products.Domain{productData},nil).Once()

		resp, err := productUsecase.GetAll()

		assert.Nil(t, err)
		assert.Contains(t, resp, productData)
	})
	t.Run("Valid Test | No Content", func(t *testing.T){
		mockMembershipProductRepo.On("GetAll").
			Return(nil,assert.AnError).Once()

		resp, err := productUsecase.GetAll()

		assert.NotNil(t, err)
		assert.Equal(t, []membership_products.Domain{}, resp)
	})
}

func TestGetByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMembershipProductRepo.On("GetByID", mock.AnythingOfType("uint")).
			Return(productData, nil).Once()

		resp, err := productUsecase.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, productData, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockMembershipProductRepo.On("GetByID", mock.AnythingOfType("uint")).
			Return(membership_products.Domain{}, assert.AnError).Once()

		resp, err := productUsecase.GetByID(1)

		assert.NotNil(t, err)
		assert.Equal(t, membership_products.Domain{}, resp)
	})
	t.Run("Invalid Test | Record Not Found", func(t *testing.T){
		mockMembershipProductRepo.On("GetByID", mock.AnythingOfType("uint")).
			Return(membership_products.Domain{}, gorm.ErrRecordNotFound).Once()

		resp, err := productUsecase.GetByID(1)

		assert.NotNil(t, err)
		assert.Equal(t, business.ErrProductNotFound, err)
		assert.Equal(t, membership_products.Domain{}, resp)
	})
}

func TestUpdateByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMembershipProductRepo.On("UpdateByID", mock.AnythingOfType("uint"), mock.Anything).
			Return(productData, nil).Once()

		resp, err := productUsecase.UpdateByID(1, &productData)

		assert.Nil(t, err)
		assert.Equal(t, productData, resp)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockMembershipProductRepo.On("UpdateByID", mock.AnythingOfType("uint"), mock.Anything).
			Return(membership_products.Domain{}, assert.AnError).Once()

		resp, err := productUsecase.UpdateByID(1, &productData)

		assert.NotNil(t, err)
		assert.Equal(t, membership_products.Domain{}, resp)
	})
	t.Run("Invalid Test | Record Not Found", func(t *testing.T){
		mockMembershipProductRepo.On("UpdateByID", mock.AnythingOfType("uint"), mock.Anything).
			Return(membership_products.Domain{}, gorm.ErrRecordNotFound).Once()

		resp, err := productUsecase.UpdateByID(1, &productData)

		assert.NotNil(t, err)
		assert.Equal(t, business.ErrProductNotFound, err)
		assert.Equal(t, membership_products.Domain{}, resp)
	})
}

func TestDeleteByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockMembershipProductRepo.On("DeleteByID", mock.AnythingOfType("uint")).
			Return(nil).Once()

		err := productUsecase.DeleteByID(1)

		assert.Nil(t, err)
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		mockMembershipProductRepo.On("DeleteByID", mock.AnythingOfType("uint")).
			Return(assert.AnError).Once()

		err := productUsecase.DeleteByID(1)

		assert.NotNil(t, err)
	})
	t.Run("Invalid Test | Record Not Found", func(t *testing.T){
		mockMembershipProductRepo.On("DeleteByID", mock.AnythingOfType("uint")).
			Return(gorm.ErrRecordNotFound).Once()

		err := productUsecase.DeleteByID(1)

		assert.NotNil(t, err)
		assert.Equal(t, business.ErrProductNotFound, err)
	})
}