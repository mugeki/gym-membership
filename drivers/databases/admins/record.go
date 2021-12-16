package admins

import (
	"gym-membership/business/admins"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admins struct {
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

func (rec *Admins) toDomain() admins.Domain {
	return admins.Domain{
		ID:        rec.ID,
		UUID:      rec.UUID,
		Username:  rec.Username,
		Password:  rec.Password,
		Email:     rec.Email,
		FullName:  rec.FullName,
		Gender:    rec.Gender,
		Telephone: rec.Telephone,
		Address:   rec.Address,
		CreatedAt: rec.CreatedAt,
	}
}

func fromDomain(domain admins.Domain) *Admins {
	return &Admins{
		Model: gorm.Model{
			ID:        domain.ID,
			CreatedAt: domain.CreatedAt,
		},
		UUID:      domain.UUID,
		Username:  domain.Username,
		Password:  domain.Password,
		Email:     domain.Email,
		FullName:  domain.FullName,
		Gender:    domain.Gender,
		Telephone: domain.Telephone,
		Address:   domain.Address,
	}
}
