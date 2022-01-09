package transactionClass

import (
	"gym-membership/drivers/databases/admins"
	"gym-membership/drivers/databases/class"
	"gym-membership/drivers/databases/users"

	"gorm.io/gorm"
)

type TransactionClass struct {
	gorm.Model
	ID      uint
	UserID  uint
	User   users.Users `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	AdminID uint
	Admin	admins.Admins `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Status  string
	ClassID int
	Class   class.Class `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}
