package admins

import (
	"gorm.io/gorm"
)

type Admins struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Password  string
	Email     string
	FullName  string
	Gender    string
	Telephone string
	Address   string
	IsSuperAdmin bool
}
