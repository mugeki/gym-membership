package class

import (
	"time"

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
	Description     string
	AvailableStatus bool
	IsOnline        bool
	Date            string
	Location        string
	CreatedAt       time.Time
}
