package members

import (
	"errors"
	"gym-membership/business"

	"gorm.io/gorm"
)

type memberUsecase struct {
	memberRepository Repository
}

func NewMemberUsecase(memberRepo Repository) Usecase {
	return &memberUsecase{
		memberRepository: memberRepo,
	}
}

func (uc *memberUsecase) GetByUserID(userID uint) (Domain, error) {
	data, err := uc.memberRepository.GetByUserID(userID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err){
			return Domain{}, business.ErrUserNotFound
		}
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}
