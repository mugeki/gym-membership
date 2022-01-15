package class

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
	// "gym-membership/helper/encrypt"
	// "github.com/google/uuid"
)

type classUsecase struct {
	classRepository Repository
	jwtAuth         *middleware.ConfigJWT
}

func NewClassUsecase(classRepo Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &classUsecase{
		classRepository: classRepo,
		jwtAuth:         jwtauth,
	}
}

func (uc *classUsecase) Insert(classData *Domain) (string, error) {
	println("bussines classes", classData.Name)
	_, err := uc.classRepository.Insert(classData)
	if err != nil {
		return "", business.ErrDuplicateData
	}

	return "", nil
}

func (uc *classUsecase) GetAll(name string, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 10
	}
	res, totalData, err := uc.classRepository.GetAll(name, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	return res, offset, limit, totalData, nil
}

func (uc *classUsecase) UpdateClassByID(id uint, classData *Domain) (Domain, error) {
	dataUpdated, err := uc.classRepository.UpdateClassByID(id, classData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return dataUpdated, nil
}

func (uc *classUsecase) UpdateParticipant(idClass int) (string, error) {
	_, err := uc.classRepository.UpdateParticipant(uint(idClass))
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}

// func (uc *classUsecase) ScheduleByID(idUser uint) ([]Domain, error) {
// 	_, err := uc.classRepository.ScheduleByID(idUser)
// 	if err != nil {
// 		return []Domain{}, business.ErrInternalServer
// 	}
// 	return []Domain{}, nil
// }
