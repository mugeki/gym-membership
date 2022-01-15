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
	GetAll(title string, page int) ([]Domain, int, int, int64, error)
	Insert(videoData *Domain) (string, error)
	UpdateByID(id uint, videoData *Domain) (string, error)
	DeleteByID(id uint) (error)
}

type Repository interface {
	GetAll(title string, offset, limit int) ([]Domain, int64, error)
	Insert(videoData *Domain) (Domain, error)
	UpdateByID(id uint, videoData *Domain) (Domain, error)
	DeleteByID(id uint) (error)
}

