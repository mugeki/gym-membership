package users

import (
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID 			uint
	UUID 		uuid.UUID
	Username	string
	Password	string
	Email		string
	FullName 	string
	Gender 		string
	Telephone 	string
	Address 	string
	CreatedAt 	time.Time
	Token		string
}

type Usecase interface {
	Register(userData *Domain) (error)
	Login(username, password string) (Domain, error)
}

type Repository interface {
	Register(userData *Domain) (Domain, error)
	GetByUsername(username string) (Domain, error)
}

