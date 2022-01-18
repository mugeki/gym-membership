package class

import (
	"time"
)

type Domain struct {
	ID uint
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
	UpdateParticipant(idClass int) (string, error)
	GetAll(title string, page int) ([]Domain, int, int, int64, error)
	// GetClassById(idClass int) (Domain, error)
	UpdateClassByID(id uint, classData *Domain) (Domain, error)
	// ScheduleByID(id uint) ([]Domain, error)
}

type Repository interface {
	Insert(classData *Domain) (Domain, error)
	UpdateClassByID(id uint, classData *Domain) (Domain, error)
	UpdateParticipant(idClass uint) (Domain, error)
	GetAll(title string, offset, limit int) ([]Domain, int64, error)
	// GetClassById(idClass uint) (Domain, error)
	UpdateStatus(idClass uint, status bool) (Domain, error)
	IsExist(idClass uint) (Domain, error)
	// ScheduleByID(id uint) ([]Domain, error)
}
