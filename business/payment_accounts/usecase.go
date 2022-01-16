package payment_accounts

import (
	"gym-membership/business"

	"github.com/jinzhu/copier"
)

type paymentAccountUsecase struct {
	paymentAccountRepository Repository
}

func NewPaymentAccountUsecase(paymentAccountRepo Repository) Usecase {
	return &paymentAccountUsecase{
		paymentAccountRepository: paymentAccountRepo,
	}
}

func (uc *paymentAccountUsecase) GetAll() ([]Domain, error) {
	res, err := uc.paymentAccountRepository.GetAll()
	if err != nil {
		return res, business.ErrInternalServer
	}
	domain := []Domain{}
	copier.Copy(&domain, &res)
	return domain, nil
}

func (uc *paymentAccountUsecase) Insert(paymentAccount *Domain) (Domain, error) {
	res, err := uc.paymentAccountRepository.Insert(paymentAccount)
	if err != nil {
		return res, business.ErrInternalServer
	}
	domain := Domain{}
	copier.Copy(&domain, &res)
	return domain, nil
}
