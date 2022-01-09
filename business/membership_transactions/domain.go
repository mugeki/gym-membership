package membership_transactions

import (
	"time"
)

type Domain struct {
	ID uint
	UserID  uint
	AdminID uint
	Status  string
	Nominal int
	MembershipProductID uint
	Date    time.Time
}

type Usecase interface {
	Insert(membershipTransactionData *Domain) (Domain, error)
	UpdateStatus(id, idAdmin uint, status string) (error)
	GetAll(status string, idUser uint, page int) ([]Domain, int, int, int64, error)
}

type Repository interface {
	Insert(membershipTransactionData *Domain) (Domain, error)
	UpdateStatus(iid, idAdmin uint, status string) (Domain, error)
	GetAll(status string, idUser uint, offset, limit int) ([]Domain, int64, error)
}
