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

func NewMembersUsecase(membersRepository Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &membersUsecase{
		membersRepository: membersRepository,
		jwtAuth:         jwtauth,
	}
}

func (uc *membersUsecase) Insert(membersData *Domain) (string, error) {
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
