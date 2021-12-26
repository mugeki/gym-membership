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

	return "item created", nil
}

func (uc *classUsecase) UpdateKuota(idClass int) (string, error) {
	_, err := uc.classRepository.UpdateKuota(idClass)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}
