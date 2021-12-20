package classifications

import (
	"gym-membership/business/classification"

	"gorm.io/gorm"
)

type Classification struct {
	gorm.Model
	ID   uint
	Name string
}

func (rec *Classification) toDomain() classification.Domain {
	return classification.Domain{
		ID:   rec.ID,
		Name: rec.Name,
	}
}

func fromDomain(domain classification.Domain) *Classification {
	return &Classification{
		ID:   domain.ID,
		Name: domain.Name,
	}
}
