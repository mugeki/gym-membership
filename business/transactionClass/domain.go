package transactionClass

import (
	"time"
)

type Domain struct {
	ID uint
	// UUID      uuid.UUID
	UserID   uint
	AdminID  uint
	Status   string
	Nominal  int
	ClassID  int
	Location string
	Date     time.Time
}

type Usecase interface {
	Insert(classData *Domain) (Domain, error)
	UpdateStatus(id uint, status string) (string, error)
	GetAll(status string, idUser uint, page int) ([]Domain, int, int, int64, error)
	// GetActiveClass(idUser uint) ([]class.Domain, error)
}

type Repository interface {
	Insert(classData *Domain) (Domain, error)
	UpdateStatus(id uint, status string) (Domain, error)
	GetAll(status string, idUser uint, offset, limit int) ([]Domain, int64, error)
	// GetActiveClass(idUser uint) ([]class.Domain, error)
}
