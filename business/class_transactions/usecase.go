package class_transactions

import (
	"gym-membership/business"
	"gym-membership/business/class"
	"strings"
)

type classTransactionUsecase struct {
	classTransactionRepository Repository
	classRepository            class.Repository
}

func NewClassTransactionUsecase(classTransactionRepo Repository, classRepository class.Repository) Usecase {
	return &classTransactionUsecase{
		classTransactionRepository: classTransactionRepo,
		classRepository:            classRepository,
	}
}

func (uc *classTransactionUsecase) Insert(classTransactionData *Domain) (Domain, error) {
	classTransactionData.Status = "waiting for payment"
	data, err := uc.classTransactionRepository.Insert(classTransactionData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	idClass := classTransactionData.ClassID
	_, errUpdateKuota := uc.classRepository.UpdateParticipant(uint(idClass))

	if errUpdateKuota != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *classTransactionUsecase) GetAll(status string, idUser uint, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 10
	}

	res, totalData, err := uc.classTransactionRepository.GetAll(status, idUser, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	return res, offset, limit, totalData, nil
}

func (uc *classTransactionUsecase) GetActiveClass(idUser uint) ([]class.Domain, error) {
	res, err := uc.classTransactionRepository.GetActiveClass(idUser)
	if err != nil {
		return []class.Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (uc *classTransactionUsecase) UpdateStatus(id, idAdmin uint, status string) (string, error) {
	formattedStatus := strings.ReplaceAll(status, "-", " ")
	_, err := uc.classTransactionRepository.UpdateStatus(id, idAdmin, formattedStatus)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}
