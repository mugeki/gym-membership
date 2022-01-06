package trainers

import (
	"gym-membership/business"
)

// "gym-membership/helper/encrypt"
// "github.com/google/uuid"

type trainerUsecase struct {
	trainerRepository Repository
}

func NewTrainerUsecase(trainerRepo Repository) Usecase {
	return &trainerUsecase{
		trainerRepository: trainerRepo,
	}
}

func (uc *trainerUsecase) GetAll() ([]Domain, error) {
	data, err := uc.trainerRepository.GetAll()
	if err != nil {
		return []Domain{}, business.ErrInternalServer
	}
	return data, nil
}
