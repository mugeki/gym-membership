package class_transactions

import (
	"gym-membership/business/class"
	"time"
)

type Domain struct {
	ID 			uint
	UserID  	uint
	AdminID 	uint
	Status  	string
	Nominal 	int
	ClassID 	int
	CreatedAt   time.Time
}

type Usecase interface {
	Insert(classTransactioData *Domain) (Domain, error)
	UpdateStatus(idTransactionClass, idAdmin uint, status string) (string, error)
	GetAll(status string, idUser uint, page int) ([]Domain, int, int, int64, error)
	GetActiveClass(idUser uint) ([]class.Domain, error)
}

type Repository interface {
	Insert(classTransactioData *Domain) (Domain, error)
	UpdateStatus(idTransactionClass, idAdmin uint, status string) (Domain, error)
	GetAll(status string, idUser uint, offset, limit int) ([]Domain, int64, error)
	GetActiveClass(idUser uint) ([]class.Domain, error)
}
