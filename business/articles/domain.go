package articles

import (
	"time"
)

type Domain struct {
	ID                 uint
	Title              string
	ClassificationID   uint
	ClassificationName string
	AdminID            uint
	MemberOnly         bool
	UrlImage           string
	CreatedAt          time.Time
	Text               string
}

type Usecase interface {
	GetAll(title string, page int) ([]Domain, int, int, int64, error)
	Insert(articleData *Domain, adminID uint) (string, error)
	UpdateArticleByID(id uint, articleData *Domain, adminID uint) (string, error)
}

type Repository interface {
	GetAll(title string, offset, limit int) ([]Domain, int64, error)
	// GetClassificationID(classification string) (int, error)
	Insert(videoData *Domain) (Domain, error)
	UpdateByID(id uint, articleData *Domain) (Domain, error)
}
