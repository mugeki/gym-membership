package membership_transactions

import (
	"gym-membership/drivers/databases/admins"
	"gym-membership/drivers/databases/membership_products"
	"gym-membership/drivers/databases/users"

	"gorm.io/gorm"
)

type MembershipTransactions struct {
	gorm.Model
	UserID  uint
	User   users.Users `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	AdminID uint
	Admin	admins.Admins `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Status  string
	MembershipProductID uint
	MembershipProduct   membership_products.MembershipProducts `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}