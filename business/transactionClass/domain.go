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
	Nominal  string
	ClassID  int
	Location string
	Date     time.Time
}

type Usecase interface {
	Insert(classData *Domain) (string, error)
	UpdateStatus(id uint, status string) (string, error)
	GetAll(page int) ([]Domain, int, int, int64, error)
}

type Repository interface {
	Insert(classData *Domain) (Domain, error)
	UpdateStatus(id uint, status string) (Domain, error)
	GetAll(offset, limit int) ([]Domain, int64, error)
	// IsExist(idClass int) (Domain, error)
}
