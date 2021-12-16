package videos

import (
	"gym-membership/drivers/databases/admins"
	"gym-membership/drivers/databases/classifications"

	"gorm.io/gorm"
)

type Videos struct {
	gorm.Model
	Title				string
	ClassificationID	int
	Classification		classifications.Classifications	`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	AdminID				uint
	Admin				admins.Admins					`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	MemberOnly			bool
	Url					string				
}