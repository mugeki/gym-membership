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
	UpdateParticipant(idClass int) (string, error)
	GetAll(title string, page int) ([]Domain, int, int, int64, error)
	UpdateClassByID(id uint, classData *Domain) (Domain, error)
	// ScheduleByID(id uint) ([]Domain, error)
}

type Repository interface {
	Insert(classData *Domain) (Domain, error)
	UpdateClassByID(id uint, classData *Domain) (Domain, error)
	UpdateParticipant(idClass int) (Domain, error)
	GetAll(title string, offset, limit int) ([]Domain, int64, error)
	UpdateStatus(idClass int, status bool) (Domain, error)
	IsExist(idClass int) (Domain, error)
	// ScheduleByID(id uint) ([]Domain, error)
}
