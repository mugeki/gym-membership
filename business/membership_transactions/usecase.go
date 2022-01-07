package membership_transactions

import (
	"errors"
	"gym-membership/business"
	"gym-membership/business/members"
	"gym-membership/business/membership_products"
	"strings"
	"time"

	"gorm.io/gorm"
)

type membershipTransactionUsecase struct {
	membershipTransactionRepository Repository
	membershipProductRepository membership_products.Repository
	memberRepository            members.Repository
}

func NewTransactionClassUsecase(membershipTransactionRepo Repository, membershipProductRepo membership_products.Repository, memberRepo members.Repository) Usecase {
	return &membershipTransactionUsecase{
		membershipTransactionRepository: membershipTransactionRepo,
		membershipProductRepository: membershipProductRepo,
		memberRepository:            memberRepo,
	}
}

func (uc *membershipTransactionUsecase) Insert(membershipTransactionData *Domain) (Domain, error) {
	membershipTransactionData.Status = "waiting for payment"
	data, err := uc.membershipTransactionRepository.Insert(membershipTransactionData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *membershipTransactionUsecase) GetAll(status string, idUser uint, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 10
	}

	resStatus := strings.ReplaceAll(status, "-", " ")
	res, totalData, err := uc.membershipTransactionRepository.GetAll(resStatus, idUser, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	return res, offset, limit, totalData, nil
}

func (uc *membershipTransactionUsecase) UpdateStatus(id uint, status string) (error) {
	formattedStatus := strings.ReplaceAll(status, "-", " ")
	data, err := uc.membershipTransactionRepository.UpdateStatus(id, formattedStatus)
	if err != nil {
		return business.ErrInternalServer
	}

	dataProduct, err := uc.membershipProductRepository.GetByID(data.MembershiProductID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err){
			return business.ErrProductNotFound
		}
		return business.ErrInternalServer
	}
	
	if status == "completed" {
		timePeriod := dataProduct.PeriodTime
		expireDate := time.Now().Add(time.Hour * 24 * time.Duration(timePeriod))
		dataMember := members.Domain{
			UserID     : data.UserID,
			ExpireDate : expireDate,
		}
		err = uc.memberRepository.Insert(&dataMember)
	}

	if err != nil {
		return business.ErrInternalServer
	}
	return nil
}
