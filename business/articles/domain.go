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
	GetAll() ([]Domain, error)
	Insert(articleData *Domain, adminID uint) (string, error)
	UpdateArticleByID(id uint, articleData *Domain, adminID uint) (string, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
	// GetClassificationID(classification string) (int, error)
	Insert(videoData *Domain) (Domain, error)
	UpdateByID(id uint, articleData *Domain) (Domain, error)
}
