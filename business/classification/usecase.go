package classification

import "gym-membership/business"

type classificationUsecase struct {
	classificationRepository Repository
}

func NewClassificationUsecase(classificationRepo Repository) Usecase {
	return &classificationUsecase{
		classificationRepository: classificationRepo,
	}
}

func (uc *classificationUsecase) GetAll() ([]Domain, error) {
	res, err := uc.classificationRepository.GetAll()
	if err != nil {
		return res, business.ErrInternalServer
	}
	return res, nil
}

func (uc *classificationUsecase) Insert(classification *Domain) (Domain, error) {
	res, err := uc.classificationRepository.Insert(classification)
	if err != nil {
		return res, business.ErrInternalServer
	}
	return res, nil
}
