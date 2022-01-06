package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UUID 		uuid.UUID
	Username	string			`gorm:"unique"`
	Password	string
	Email		string
	FullName 	string	
	Gender 		string
	Telephone 	string	
	Address 	string
	UrlImage	string
}