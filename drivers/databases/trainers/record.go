package trainers

import (
	"gorm.io/gorm"
)

type Trainers struct {
	gorm.Model
	ID       uint
	Fullname string
	UrlImage string
}
