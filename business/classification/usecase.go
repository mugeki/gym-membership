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

func (uc *classificationUsecase) GetClassificationID(name string) (uint, error) {
	res, err := uc.classificationRepository.GetClassificationID(name)
	if err != nil {
		return res, business.ErrInternalServer
	}
	return res, nil
}
