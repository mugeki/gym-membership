package class_transactions

import (
	"gym-membership/drivers/databases/admins"
	"gym-membership/drivers/databases/class"
	"gym-membership/drivers/databases/payment_accounts"
	"gym-membership/drivers/databases/users"

	"gorm.io/gorm"
)

type ClassTransaction struct {
	gorm.Model
	ID                uint
	UserID            uint
	User              users.Users   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	AdminID           uint          `gorm:"default:1"`
	UrlImageOfReceipt string        `json:"url_image_of_receipt"`
	Admin             admins.Admins `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Status            string
	ClassID           int
	Class             class.Class `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	PaymentID         int
	Payment           payment_accounts.PaymentAccount `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}
