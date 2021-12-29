package transactionClass

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
	"gym-membership/business/class"
	"strings"
	// "gym-membership/helper/encrypt"
	// "github.com/google/uuid"
)

type transactionClassUsecase struct {
	transactionClassRepository Repository
	classRepository            class.Repository
	jwtAuth                    *middleware.ConfigJWT
}

func NewTransactionClassUsecase(transactionClassRepo Repository, classRepository class.Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &transactionClassUsecase{
		transactionClassRepository: transactionClassRepo,
		classRepository:            classRepository,
		jwtAuth:                    jwtauth,
	}
}

func (uc *transactionClassUsecase) Insert(classData *Domain) (Domain, error) {
	classData.Status = "waiting for payment"
	data, err := uc.transactionClassRepository.Insert(classData)
	if err != nil {
		return Domain{}, business.ErrDuplicateData
	}
	idClass := classData.ClassID
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

	resStatus := strings.ReplaceAll(status, "-", " ")
	res, totalData, err := uc.transactionClassRepository.GetAll(resStatus, idUser, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	return res, offset, limit, totalData, nil
}

// func (uc *transactionClassUsecase) GetActiveClass(idUser uint) ([]class.Domain, error) {
// 	res, err := uc.transactionClassRepository.GetActiveClass(idUser)
// 	if err != nil {
// 		return []class.Domain{}, business.ErrInternalServer
// 	}
// 	return res, nil
// }

func (uc *transactionClassUsecase) UpdateStatus(id uint, status string) (string, error) {
	_, err := uc.transactionClassRepository.UpdateStatus(id, status)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}