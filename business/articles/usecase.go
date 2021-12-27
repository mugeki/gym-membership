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

func (uc *articleUsecase) Insert(articleData *Domain) (string, error) {
	// classificationID, _ := uc.classificationRepository.GetClassificationID(articleData.ClassificationName)
	// articleData.ClassificationID = classificationID
	// articleData.AdminID = adminID
	// println("cek data text", articleData.Text)
	_, err := uc.articleRepository.Insert(articleData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "item created", nil
}

func (uc *articleUsecase) UpdateArticleByID(id uint, videoData *Domain) (string, error) {
	println("cek id", id)
	// classificationID, _ := uc.classificationRepository.GetClassificationID(videoData.ClassificationName)
	// videoData.ClassificationID = classificationID
	// videoData.AdminID = adminID
	_, err := uc.articleRepository.UpdateByID(id, videoData)
	if err != nil {
		return "", business.ErrInternalServer
	}
	return "item edited", nil
}
