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
	membershipProductRepository     membership_products.Repository
	memberRepository                members.Repository
}

func NewMembershipTransactionUsecase(membershipTransactionRepo Repository, membershipProductRepo membership_products.Repository, memberRepo members.Repository) Usecase {
	return &membershipTransactionUsecase{
		membershipTransactionRepository: membershipTransactionRepo,
		membershipProductRepository:     membershipProductRepo,
		memberRepository:                memberRepo,
	}
}

func (uc *membershipTransactionUsecase) Insert(membershipTransactionData *Domain) (Domain, error) {
	membershipTransactionData.Status = "waiting-for-payment"
	data, err := uc.membershipTransactionRepository.Insert(membershipTransactionData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *membershipTransactionUsecase) GetAll(date time.Time, status string, idUser uint, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 10
	}

	res, totalData, err := uc.membershipTransactionRepository.GetAll(date, status, idUser, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	for i := 0; i < len(res); i++ {
		res[i].Status = strings.ReplaceAll(res[i].Status, "-", " ")
	}
	return res, offset, limit, totalData, nil
}

func (uc *membershipTransactionUsecase) GetAllByUser(idUser uint) ([]Domain, error) {

	res, err := uc.membershipTransactionRepository.GetAllByUser(idUser)
	if err != nil {
		return []Domain{}, business.ErrInternalServer
	}
	for i := 0; i < len(res); i++ {
		res[i].Status = strings.ReplaceAll(res[i].Status, "-", " ")
	}
	return res, nil
}

func (uc *membershipTransactionUsecase) GetByID(idTransaction uint) (Domain, error) {
	res, err := uc.membershipTransactionRepository.GetByID(idTransaction)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	res.Status = strings.ReplaceAll(res.Status, "-", " ")
	return res, nil
}

func (uc *membershipTransactionUsecase) UpdateStatus(id, idAdmin uint, status string) error {
	data, err := uc.membershipTransactionRepository.UpdateStatus(id, idAdmin, status)
	if err != nil {
		return business.ErrInternalServer
	}

	dataProduct, err := uc.membershipProductRepository.GetByID(data.MembershipProductID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return business.ErrProductNotFound
		}
		return business.ErrInternalServer
	}

	if status == "accepted" {
		timePeriod := dataProduct.PeriodTime
		expireDate := time.Now().Add(time.Hour * 24 * time.Duration(timePeriod))
		dataMember := members.Domain{
			UserID:     data.UserID,
			ExpireDate: expireDate,
		}
		err = uc.memberRepository.Insert(&dataMember)
		if err != nil {
			return business.ErrInternalServer
		}
	}

	return nil
}

func (uc *membershipTransactionUsecase) UpdateReceipt(id uint, urlImage string) (string, error) {
	status := "waiting-for-confirmation"
	_, err := uc.membershipTransactionRepository.UpdateReceipt(id, urlImage, status)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}

func (uc *membershipTransactionUsecase) UpdateStatusToFailed(transactionID uint) (Domain, error) {
	status := "failed"
	data, err := uc.membershipTransactionRepository.UpdateStatusToFailed(transactionID, status)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}
