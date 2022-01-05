package articles

import (
	"gym-membership/drivers/databases/admins"
	"gym-membership/drivers/databases/classifications"

	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	Title            string
	ClassificationID uint
	Classification   classifications.Classification `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	AdminID          uint
	Admin            admins.Admins `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	MemberOnly       bool
	UrlImage         string
	Text             string
}
