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

func (uc *videoUsecase) GetAll(title string, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page-1)*10
	}
	res, totalData, err := uc.videoRepository.GetAll(title, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	return res, offset, limit, totalData, nil
}

func (uc *videoUsecase) Insert(videoData *Domain) (string, error) {
	_, err := uc.videoRepository.Insert(videoData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}

func (uc *videoUsecase) UpdateByID(id uint, videoData *Domain) (string, error) {
	_, err := uc.videoRepository.UpdateByID(id, videoData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}
