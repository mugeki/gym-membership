package request

import "gym-membership/business/videos"

type Videos struct {
	Title          	string 	`json:"title" valid:"required"`
	Classification 	string 	`json:"classification" valid:"required"`
	MemberOnly		bool	`json:"memberOnly"`
	Url            	string 	`json:"url" valid:"required, url"`
}

func (req *Videos) ToDomain() *videos.Domain {
	return &videos.Domain{		
		Title				: req.Title,
		ClassificationName	: req.Classification,
		MemberOnly			: req.MemberOnly,
		Url					: req.Url,
	}
}