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
	TrainerName     string
	TrainerImage    string
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
	GetAll(title string, page int) ([]Domain, int, int, int64, error)
	UpdateClassByID(id uint, articleData *Domain) (string, error)
}

type Repository interface {
	Insert(classData *Domain) (Domain, error)
	UpdateClassByID(id uint, articleData *Domain) (Domain, error)
	UpdateKuota(idClass int) (string, error)
	GetAll(title string, offset, limit int) ([]Domain, int64, error)
	UpdateStatusToFalse(idClass int) (string, error)
	IsExist(idClass int) (Domain, error)
}
