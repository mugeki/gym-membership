package request

import (
	"gym-membership/business/users"

	"github.com/google/uuid"
)

type Users struct {			
	ID 			uint
	UUID 		uuid.UUID
	Username	string		`json:"username" valid:"required,minstringlength(6)"`
	Password	string		`json:"password" valid:"required,minstringlength(6)"`
	Email		string		`json:"email" valid:"required,email"`
	FullName 	string		`json:"fullname" valid:"required"`
	Gender 		string		`json:"gender" valid:"required"`
	Telephone 	string		`json:"telephone" valid:"required,numeric"`
	Address 	string		`json:"address" valid:"-"`
}

type UsersLogin struct{
	Username    string	`json:"username" valid:"required,minstringlength(6)"`
	Password    string	`json:"password" valid:"required,minstringlength(6)"`
}

func (req *Users) ToDomain() (*users.Domain) {
	return &users.Domain{
		ID 			: req.ID,
		UUID 		: req.UUID,
		Username	: req.Username,
		Password	: req.Password,
		Email		: req.Email,
		FullName 	: req.FullName,
		Gender 		: req.Gender,
		Telephone 	: req.Telephone,
		Address 	: req.Address,
	}
}