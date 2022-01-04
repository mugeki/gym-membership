package request

import (
	"gym-membership/business/classification"
)

type Classification struct {
	ID                 uint   `json:"id"`
	ClassificationName string `json:"classificationName"`
}

func (req *Classification) ToDomain() *classification.Domain {
	return &classification.Domain{
		ID:   req.ID,
		Name: req.ClassificationName,
	}
}

// type GetAll struct {
// 	Username string `json:"username" valid:"required,minstringlength(6)"`
// 	Password string `json:"password" valid:"required,minstringlength(6)"`
// }
