package articles

import (
	"gym-membership/business"
	classification "gym-membership/business/classification"
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

func (uc *articleUsecase) GetAll() ([]Domain, error) {
	res, err := uc.articleRepository.GetAll()
	if err != nil {
		return []Domain{}, business.ErrInternalServer
	}
	return res, nil
}

func (uc *articleUsecase) Insert(articleData *Domain, adminID uint) (string, error) {
	classificationID, _ := uc.classificationRepository.GetClassificationID(articleData.ClassificationName)
	articleData.ClassificationID = classificationID
	articleData.AdminID = adminID
	// println("cek data text", articleData.Text)
	_, err := uc.articleRepository.Insert(articleData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "item created", nil
}

func (uc *articleUsecase) UpdateArticleByID(id uint, videoData *Domain, adminID uint) (string, error) {
	println("cek id", id)
	classificationID, _ := uc.classificationRepository.GetClassificationID(videoData.ClassificationName)
	videoData.ClassificationID = classificationID
	videoData.AdminID = adminID
	_, err := uc.articleRepository.UpdateByID(id, videoData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "item edited", nil
}
