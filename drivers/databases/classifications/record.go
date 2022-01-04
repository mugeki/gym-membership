package classifications

import (
	"gorm.io/gorm"
)

type Classification struct {
	gorm.Model
	ID   uint
	Name string
}
