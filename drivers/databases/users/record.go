package users

import (
	"gym-membership/business/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UUID 		uuid.UUID
	Username	string		`gorm:"unique"`
	Password	string
	Email		string
	FullName 	string	
	Gender 		users.Gender
	Telephone 	string	
	Address 	string	
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		ID        	: rec.ID,
		UUID      	: rec.UUID,
		Username	: rec.Username,
		Password	: rec.Password,
		Email		: rec.Email,
		FullName  	: rec.FullName,
		Gender    	: rec.Gender,
		Telephone 	: rec.Telephone,
		Address   	: rec.Address,
		CreatedAt 	: rec.CreatedAt,
	}
}

func fromDomain(domain users.Domain) *Users {
	return &Users{
		Model		: gorm.Model {
						ID			: domain.ID,
						CreatedAt	: domain.CreatedAt,
					},
		UUID      	: domain.UUID,
		Username	: domain.Username,
		Password	: domain.Password,
		Email		: domain.Email,
		FullName  	: domain.FullName,
		Gender    	: domain.Gender,
		Telephone 	: domain.Telephone,
		Address   	: domain.Address,
	}
}