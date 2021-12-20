package videos

import (
	"time"
)

type Domain struct {
	ID					uint
	Title				string
	ClassificationID	int
	ClassificationName	string
	AdminID				uint
	MemberOnly			bool
	Url					string
	CreatedAt			time.Time
}

type Usecase interface {
	GetAll(title string, page int) ([]Domain, error)
	Insert(videoData *Domain, adminID uint) (string, error)
	UpdateByID(id uint, videoData *Domain, adminID uint) (string, error)
}

type Repository interface {
	GetAll(title string, offset, limit int) ([]Domain, error)
	GetClassificationID(classification string) (int, error)
	Insert(videoData *Domain) (Domain, error)
	UpdateByID(id uint, videoData *Domain) (Domain, error)
}

