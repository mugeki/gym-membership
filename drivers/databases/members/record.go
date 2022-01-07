package members

import (
	"gym-membership/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Members struct {
	gorm.Model
	UserID     uint
	Users      users.Users `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	ExpireDate time.Time
}