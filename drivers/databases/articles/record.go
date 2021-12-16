package articles

import (
	"gym-membership/business/articles"
	"gym-membership/drivers/databases/admins"

	"gorm.io/gorm"
)

// type VideoClassifications struct {
// 	ID		int		`gorm:"primarykey"`
// 	Name	string
// }

type Articles struct {
	gorm.Model
	Title            string
	ClassificationID uint
	// Classification		VideoClassifications	`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	AdminID    uint
	Admin      admins.Admins `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	MemberOnly bool
	UrlImage   string
	Text       string
}

func (rec *Articles) toDomain() articles.Domain {
	return articles.Domain{
		ID:               rec.ID,
		Title:            rec.Title,
		ClassificationID: rec.ClassificationID,
		// ClassificationName	: rec.Classification.Name,
		AdminID:    rec.AdminID,
		MemberOnly: rec.MemberOnly,
		UrlImage:   rec.UrlImage,
		CreatedAt:  rec.CreatedAt,
		Text:       rec.Text,
	}
}

func toDomainArray(rec []Articles) []articles.Domain {
	arr := []articles.Domain{}
	for _, val := range rec {
		arr = append(arr, val.toDomain())
	}
	return arr
}

func fromDomain(domain articles.Domain) *Articles {
	return &Articles{
		Model: gorm.Model{
			ID:        domain.ID,
			CreatedAt: domain.CreatedAt,
		},
		Title:            domain.Title,
		ClassificationID: domain.ClassificationID,
		AdminID:          domain.AdminID,
		MemberOnly:       domain.MemberOnly,
		UrlImage:         domain.UrlImage,
	}
}
