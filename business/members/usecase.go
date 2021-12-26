package members

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
	// "gym-membership/helper/encrypt"
	// "github.com/google/uuid"
)

type membersUsecase struct {
	membersRepository Repository
	jwtAuth         *middleware.ConfigJWT
}

func NewMembersUsecase(membersRepo Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &membersUsecase{
		membersRepository: membersRepo,
		jwtAuth:         jwtauth,
	}
}

func (uc *membersUsecase) Insert(membersData *Domain) (string, error) {
	println("bussines members", membersData.Name)
	_, err := uc.membersRepository.Insert(membersData)
	if err != nil {
		return "", business.ErrDuplicateData
	}

	return "item created", nil
}

func (uc *membersUsecase) GetByUserID(idMembers int) (string, error) {
	_, err := uc.MembersRepository.GetByUserID(idMembers)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}
