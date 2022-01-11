package transactionClass

import (
	"gym-membership/business"
	"gym-membership/business/class"
	// "gym-membership/helper/encrypt"
	// "github.com/google/uuid"
)

type transactionClassUsecase struct {
	transactionClassRepository Repository
	classRepository            class.Repository
	// jwtAuth                    *middleware.ConfigJWT
}

func NewTransactionClassUsecase(transactionClassRepo Repository, classRepository class.Repository) Usecase {
	return &transactionClassUsecase{
		transactionClassRepository: transactionClassRepo,
		classRepository:            classRepository,
		// jwtAuth:                    jwtauth,
	}
}

func (uc *transactionClassUsecase) Insert(transactionClassData *Domain) (Domain, error) {
	transactionClassData.Status = "waiting-for-payment"
	data, err := uc.transactionClassRepository.Insert(transactionClassData)
	if err != nil {
		return Domain{}, business.ErrDuplicateData
	}
	idClass := transactionClassData.ClassID
	_, errUpdateKuota := uc.classRepository.UpdateParticipant(idClass)

	if errUpdateKuota != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *transactionClassUsecase) GetAll(status string, idUser uint, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 10
	}

	res, totalData, err := uc.transactionClassRepository.GetAll(status, idUser, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	return res, offset, limit, totalData, nil
}

func (uc *transactionClassUsecase) GetActiveClass(idUser uint) ([]class.Domain, error) {
	res, err := uc.transactionClassRepository.GetActiveClass(idUser)
	if err != nil {
		return []class.Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (uc *transactionClassUsecase) UpdateStatus(id, idAdmin uint, status string) (string, error) {
	_, err := uc.transactionClassRepository.UpdateStatus(id, idAdmin, status)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}
