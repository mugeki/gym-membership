package request

import (
	"gym-membership/business/articles"
)

type Articles struct {
	ID                 uint   `json:"id"`
	Title              string `json:"title"`
	ClassificationName string `json:"classificationName"`
	AdminID            uint   `json:"adminId"`
	MemberOnly         bool   `json:"memberOnly"`
	UrlImage           string `json:"urlImage"`
	Text               string `json:"text"`
}

func (req *Articles) ToDomain() *articles.Domain {
	return &articles.Domain{
		Title:              req.Title,
		ClassificationName: req.ClassificationName,
		MemberOnly:         req.MemberOnly,
		UrlImage:           req.UrlImage,
		Text:               req.Text,
	}
}

// type GetAll struct {
// 	Username string `json:"username" valid:"required,minstringlength(6)"`
// 	Password string `json:"password" valid:"required,minstringlength(6)"`
// }
