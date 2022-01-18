package payment_accounts

import (
	"gorm.io/gorm"
)

type PaymentAccount struct {
	gorm.Model
	Name      string
	NoCard    string
	OwnerName string
	Desc      string
}
