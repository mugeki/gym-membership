package class

import (
	"time"

	"gym-membership/drivers/databases/trainers"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name            string
	UrlImage        string
	Price           int
	Kuota           int
	Participant     int
	TrainerId       int
	Trainers        trainers.Trainers `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;foreignKey:TrainerId"`
	Description     string
	AvailableStatus bool
	IsOnline        bool
	Date            string
	Location        string
	CreatedAt       time.Time
}
