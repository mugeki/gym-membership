package articles

import (
	"gym-membership/business/videos"
	"gym-membership/drivers/databases/admins"

	"gorm.io/gorm"
)

type VideoClassifications struct {
	ID   int `gorm:"primarykey"`
	Name string
}

type Videos struct {
	gorm.Model
	Title            string
	ClassificationID int
	Classification   VideoClassifications `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	AdminID          uint
	Admin            admins.Admins `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	MemberOnly       bool
	Url              string
}

func (rec *Videos) toDomain() videos.Domain {
	return videos.Domain{
		ID:                 rec.ID,
		Title:              rec.Title,
		ClassificationID:   rec.ClassificationID,
		ClassificationName: rec.Classification.Name,
		AdminID:            rec.AdminID,
		MemberOnly:         rec.MemberOnly,
		Url:                rec.Url,
		CreatedAt:          rec.CreatedAt,
	}
}

func toDomainArray(rec []Videos) []videos.Domain {
	arr := []videos.Domain{}
	for _, val := range rec {
		arr = append(arr, val.toDomain())
	}
	return arr
}

func fromDomain(domain videos.Domain) *Videos {
	return &Videos{
		Model: gorm.Model{
			ID:        domain.ID,
			CreatedAt: domain.CreatedAt,
		},
		Title:            domain.Title,
		ClassificationID: domain.ClassificationID,
		AdminID:          domain.AdminID,
		MemberOnly:       domain.MemberOnly,
		Url:              domain.Url,
	}
}
