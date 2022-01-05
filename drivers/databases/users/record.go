package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uint
	UUID      uuid.UUID
	Username  string `gorm:"unique"`
	Password  string
	Email     string
	FullName  string
	Gender    string
	Telephone string
	Address   string
}