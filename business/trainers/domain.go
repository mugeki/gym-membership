package trainers

import (
	"time"
)

type Domain struct {
	ID          uint
	Fullname    string
	UrlImage    string
	CreatedAt   time.Time
	DeletededAt time.Time
}

type Usecase interface {
	GetAll() ([]Domain, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
}
