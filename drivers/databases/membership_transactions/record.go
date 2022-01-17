package membership_transactions

import (
	"gym-membership/drivers/databases/admins"
	"gym-membership/drivers/databases/membership_products"
	"gym-membership/drivers/databases/payment_accounts"
	"gym-membership/drivers/databases/users"

	"gorm.io/gorm"
)

type MembershipTransactions struct {
	gorm.Model
	UserID              uint
	User                users.Users   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	UrlImageOfReceipt   string        `json:"url_image_of_receipt"`
	AdminID             uint          `gorm:"default:1"`
	Admin               admins.Admins `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Status              string
	MembershipProductID uint
	MembershipProduct   membership_products.MembershipProducts `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	PaymentID           int
	Payment             payment_accounts.PaymentAccount `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}
