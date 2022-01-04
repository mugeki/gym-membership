package response

import (
	"gym-membership/business/classification"
)

type Classification struct {
	ID                 uint   `json:"id"`
	ClassificationName string `json:"classificationName"`
}

func (req *Classification) FromDomain() *classification.Domain {
	return &classification.Domain{
		ID:   req.ID,
		Name: req.ClassificationName,
	}
}
