package membership_products

import (
	"errors"
	"gym-membership/business"

	"gorm.io/gorm"
)

type membershipProductsUsecase struct {
	membershipProductsRepository Repository
}

func NewMembershipProductsUsecase(membershipProductsRepository Repository) Usecase {
	return &membershipProductsUsecase{
		membershipProductsRepository: membershipProductsRepository,
	}
}

func (uc *membershipProductsUsecase) Insert(newData *Domain)  error {
	err := uc.membershipProductsRepository.Insert(newData)
	if err != nil {
		return business.ErrInternalServer
	}

	return nil
}

func (uc *membershipProductsUsecase) GetAll() ([]Domain, error) {
	res, err := uc.membershipProductsRepository.GetAll()
	if err != nil {
		return []Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (uc *membershipProductsUsecase) GetByID(id uint) (Domain, error) { 
	res, err := uc.membershipProductsRepository.GetByID(id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return Domain{}, business.ErrProductNotFound
		}
		return Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (uc *membershipProductsUsecase) UpdateByID(id uint, newData *Domain) error{
	err := uc.membershipProductsRepository.UpdateByID(id,newData)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err){
			return business.ErrProductNotFound
		} else {
			return business.ErrInternalServer
		}
	}
	return nil
}

func (uc *membershipProductsUsecase) DeleteByID(id uint) error{
	err := uc.membershipProductsRepository.DeleteByID(id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err){
			return business.ErrProductNotFound
		} else {
			return business.ErrInternalServer
		}
	}
	return nil
}

