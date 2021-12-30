package admins

import "gorm.io/gorm"

type Admins struct {
	gorm.Model
	Username		string
	Password		string
	FullName		string
	Gender			string
	Telephone		string
	Address			string
}