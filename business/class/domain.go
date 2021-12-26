package class

import (
	"time"
)

type Domain struct {
	ID uint
	// UUID      uuid.UUID
	Name            string
	UrlImage        string
	Price           int
	Kuota           int
	Participant     int
	TrainerId       int
	Description     string
	AvailableStatus bool
	IsOnline        bool
	Date            string
	Location        string
	CreatedAt       time.Time
}

type Usecase interface {
	Insert(classData *Domain) (string, error)
	UpdateKuota(idClass int) (string, error)
}

type Repository interface {
	Insert(classData *Domain) (Domain, error)
	UpdateKuota(idClass int) (string, error)
	UpdateStatus(idClass int) (string, error)
}
