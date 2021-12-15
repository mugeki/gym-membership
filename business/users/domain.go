package users

import (
	"time"

	"github.com/google/uuid"
)

type Gender string
const (
	Male Gender = "Male"
	Female Gender = "Female"
)

type Domain struct {
	ID 			uint
	UUID 		uuid.UUID
	Username	string
	Password	string
	Email		string
	FullName 	string
	Gender 		Gender
	Telephone 	string
	Address 	string
	CreatedAt 	time.Time
}

type Usecase interface {
	Register(userData *Domain) (string, error)
	Login(username, password string) (string, error)
}

type Repository interface {
	Register(userData *Domain) (Domain, error)
	GetByUsername(username string) (Domain, error)
}

