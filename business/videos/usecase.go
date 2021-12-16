package videos

import (
	"gym-membership/business"
)

type videoUsecase struct {
	videoRepository Repository
}

func NewVideoUsecase(videoRepo Repository) Usecase {
	return &videoUsecase{
		videoRepository: videoRepo,
	}
}

func (uc *videoUsecase) GetAll() ([]Domain, error) {
	res, err := uc.videoRepository.GetAll()
	if err != nil {
		return []Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (uc *videoUsecase) Insert(videoData *Domain, adminID uint) (string, error) {
	classificationID, _ := uc.videoRepository.GetClassificationID(videoData.ClassificationName)
	videoData.ClassificationID = classificationID
	videoData.AdminID = adminID
	_, err := uc.videoRepository.Insert(videoData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "item created", nil
}

func (uc *videoUsecase) UpdateByID(id uint, videoData *Domain, adminID uint) (string, error) {
	classificationID, _ := uc.videoRepository.GetClassificationID(videoData.ClassificationName)
	videoData.ClassificationID = classificationID
	videoData.AdminID = adminID
	_, err := uc.videoRepository.UpdateByID(id, videoData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "item edited", nil
}
