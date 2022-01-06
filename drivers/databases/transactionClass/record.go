package transactionClass

import (
	"gym-membership/drivers/databases/class"
	"gym-membership/drivers/databases/users"

	"gorm.io/gorm"
	// "gym-membership/drivers/databases/admins"
)

type TransactionClass struct {
	gorm.Model
	ID      uint
	UserID  uint
	Users   users.Users `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;foreignKey:UserID"`
	AdminID uint
	Status  string
	ClassID int
	Class   class.Class `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;foreignKey:ClassID"`
}
