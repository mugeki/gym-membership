package membership_products

import (
	"gym-membership/business"
	// "gym-membership/helper/encrypt"
	// "github.com/google/uuid"
	"fmt"
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

func (uc *membershipProductsUsecase) GetByUserID(idMembershipProducts uint) (string, error) { 
	fmt.Println("usecase ",idMembershipProducts)
	_, err := uc.membershipProductsRepository.GetByUserID(idMembershipProducts)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}
