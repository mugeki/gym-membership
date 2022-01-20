package videos

import (
	"errors"
	"gym-membership/business"

	"gorm.io/gorm"
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

func (uc *videoUsecase) GetByID(id uint) (Domain, error) {
	data, err := uc.videoRepository.GetByID(id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err){
			return Domain{}, business.ErrArticleNotFound
		}
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *videoUsecase) Insert(videoData *Domain) (Domain, error) {
	data, err := uc.videoRepository.Insert(videoData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *videoUsecase) UpdateByID(id uint, videoData *Domain) (Domain, error) {
	data, err := uc.videoRepository.UpdateByID(id, videoData)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err){
			return Domain{}, business.ErrVideoNotFound
		} else {
			return Domain{}, business.ErrInternalServer
		}
	}
	return data, nil
}

func (uc *videoUsecase) DeleteByID(id uint) (error) {
	err := uc.videoRepository.DeleteByID(id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err){
			return business.ErrVideoNotFound
		} else {
			return business.ErrInternalServer
		}
	}
	return nil
}