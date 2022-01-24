package membership_transactions

import (
	"time"
)

type Domain struct {
	ID                  uint
	UserID              uint
	UserName            string
	AdminID             uint
	Status              string
	Nominal             int
	ProductName         string
	MembershipProductID uint
	UrlImageOfReceipt   string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	PaymentID           uint
	Payment             PaymentAccount
}
type PaymentAccount struct {
	ID        uint
	Name      string
	NoCard    string
	OwnerName string
	Desc      string
}

type Usecase interface {
	Insert(membershipTransactionData *Domain) (Domain, error)
	UpdateStatus(id, idAdmin uint, status string) error
	UpdateReceipt(idTransactionClass uint, urlImage string) (string, error)
	GetAll(date time.Time, status string, idUser uint, page int) ([]Domain, int, int, int64, error)
	GetAllByUser(idUser uint) ([]Domain, error)
	GetByID(idTransaction uint) (Domain, error)
}

type Repository interface {
	Insert(membershipTransactionData *Domain) (Domain, error)
	UpdateStatus(id, idAdmin uint, status string) (Domain, error)
	UpdateReceipt(idTransactionClass uint, urlImage, status string) (Domain, error)
	GetAll(date time.Time, status string, idUser uint, offset, limit int) ([]Domain, int64, error)
	GetAllByUser(idUser uint) ([]Domain, error)
	GetByID(idTransaction uint) (Domain, error)
}
