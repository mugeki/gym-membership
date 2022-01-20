package class_transactions

import (
	"gym-membership/business/class"
	"time"
)

type Domain struct {
	ID                uint
	UserID            uint
	UserName          string
	AdminID           uint
	UrlImageOfReceipt string
	Status            string
	Nominal           int
	ClassID           int
	ClassName         string
	UpdatedAt         time.Time
	CreatedAt         time.Time
	PaymentID         uint
	Payment           PaymentAccount
}

type PaymentAccount struct {
	ID        uint
	Name      string
	NoCard    string
	OwnerName string
	Desc      string
}

type Usecase interface {
	Insert(classTransactioData *Domain) (Domain, error)
	UpdateStatus(idTransactionClass, idAdmin uint, status string) (string, error)
	UpdateReceipt(idTransactionClass uint, urlImage string) (string, error)
	GetAll(date time.Time, status string, idUser uint, page int) ([]Domain, int, int, int64, error)
	GetAllByUser(idUser uint) ([]Domain, error)
	GetActiveClass(idUser uint) ([]class.Domain, error)
	GetByID(idTransaction uint) (Domain, error)
}

type Repository interface {
	Insert(classTransactioData *Domain) (Domain, error)
	UpdateStatus(idTransactionClass, idAdmin uint, status string) (Domain, error)
	UpdateReceipt(idTransactionClass uint, urlImage string) (Domain, error)
	GetAll(date time.Time, status string, idUser uint, offset, limit int) ([]Domain, int64, error)
	GetAllByUser(idUser uint) ([]Domain, error)
	GetActiveClass(idUser uint) ([]class.Domain, error)
	GetByID(idTransaction uint) (Domain, error)
}
