package membership_transactions

import (
	"gym-membership/drivers/databases/membership_products"
	"gym-membership/drivers/databases/users"

	"gorm.io/gorm"
)

type MembershipTransactions struct {
	gorm.Model
	ID      uint
	UserID  uint
	Users   users.Users `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	AdminID uint
	Status  string
	MembershiProductID uint
	MembershipProduct   membership_products.MembershipProducts `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}