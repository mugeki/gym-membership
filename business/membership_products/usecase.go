package membership_products

import (
	"gym-membership/business"
)

type membershipProductsUsecase struct {
	membershipProductsRepository Repository
}

func NewMembershipProductsUsecase(membershipProductsRepository Repository) Usecase {
	return &membershipProductsUsecase{
		membershipProductsRepository: membershipProductsRepository,
	}
}

func (uc *membershipProductsUsecase) Insert(membershipProductsData *Domain) (string, error) {
	_, err := uc.membershipProductsRepository.Insert(membershipProductsData)
	if err != nil {
		return "", business.ErrDuplicateData
	}

	return "item created", nil
}

func (uc *membershipProductsUsecase) GetByID(idMembershipProducts uint) (Domain, error) { 
	res, err := uc.membershipProductsRepository.GetByID(idMembershipProducts)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return res, nil
}
