package class_transactions

import (
	"gym-membership/business"
	"gym-membership/business/class"
	"strings"
	"time"
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
	classTransactionData.Status = "waiting-for-payment"
	data, err := uc.classTransactionRepository.Insert(classTransactionData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	idClass := classTransactionData.ClassID
	_, errUpdateKuota := uc.classRepository.IncreaseParticipant(uint(idClass))

	if errUpdateKuota != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *classTransactionUsecase) GetAll(date time.Time, status string, idUser uint, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 10
	}

	res, totalData, err := uc.classTransactionRepository.GetAll(date, status, idUser, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	for i := 0; i < len(res); i++ {
		res[i].Status = strings.ReplaceAll(res[i].Status, "-", " ")
	}
	return res, offset, limit, totalData, nil
}

func (uc *classTransactionUsecase) GetAllByUser(idUser uint) ([]Domain, error) {

	res, err := uc.classTransactionRepository.GetAllByUser(idUser)
	if err != nil {
		return []Domain{}, business.ErrInternalServer
	}
	for i := 0; i < len(res); i++ {
		res[i].Status = strings.ReplaceAll(res[i].Status, "-", " ")
	}
	return res, nil
}

func (uc *classTransactionUsecase) GetActiveClass(idUser uint) ([]class.Domain, error) {
	res, err := uc.classTransactionRepository.GetActiveClass(idUser)
	if err != nil {
		return []class.Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (uc *classTransactionUsecase) UpdateStatus(id, idAdmin uint, status string) (string, error) {
	_, err := uc.classTransactionRepository.UpdateStatus(id, idAdmin, status)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}

func (uc *classTransactionUsecase) UpdateReceipt(id uint, urlImage string) (string, error) {
	status := "waiting-for-confirmation"
	_, err := uc.classTransactionRepository.UpdateReceipt(id, urlImage, status)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}

func (uc *classTransactionUsecase) GetByID(transactionID uint) (Domain, error) {
	res, err := uc.classTransactionRepository.GetByID(transactionID)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	res.Status = strings.ReplaceAll(res.Status, "-", " ")
	return res, nil
}

func (uc *classTransactionUsecase) UpdateStatusToFailed(transactionID uint) (Domain, error) {
	status := "failed"
	data, err := uc.classTransactionRepository.UpdateStatusToFailed(transactionID, status)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	idClass := data.ClassID
	_, error := uc.classRepository.DecreaseParticipant(uint(idClass))
	if error != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}
