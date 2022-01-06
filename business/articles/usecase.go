package articles

import (
	"errors"
	"gym-membership/business"
	classification "gym-membership/business/classification"

	"gorm.io/gorm"
)

type articleUsecase struct {
	articleRepository        Repository
	classificationRepository classification.Repository
}

func NewArticleUsecase(articleRepo Repository, classificationRepo classification.Repository) Usecase {
	return &articleUsecase{
		articleRepository:        articleRepo,
		classificationRepository: classificationRepo,
	}
}

func (uc *articleUsecase) GetAll(title string, page int) ([]Domain, int, int, int64, error) {
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 10
	}
	res, totalData, err := uc.articleRepository.GetAll(title, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, business.ErrInternalServer
	}
	return res, offset, limit, totalData, nil
}

func (uc *articleUsecase) Insert(articleData *Domain) (Domain, error) {
	// classificationID, _ := uc.classificationRepository.GetClassificationID(articleData.ClassificationName)
	// println("cek data text", articleData.Text)
	data, err := uc.articleRepository.Insert(articleData)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *articleUsecase) UpdateArticleByID(id uint, videoData *Domain) (string, error) {
	_, err := uc.articleRepository.UpdateByID(id, videoData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "", nil
}

func (uc *articleUsecase) DeleteByID(id uint) error {
	err := uc.articleRepository.DeleteByID(id)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return business.ErrArticleNotFound
		} else {
			return business.ErrInternalServer
		}
	}
	return nil
}
