package response

import "gym-membership/business/videos"

type Videos struct {
	ID             	uint   	`json:"id"`
	Title          	string 	`json:"title"`
	Classification 	string 	`json:"classification"`
	AdminID        	uint   	`json:"adminId"`
	MemberOnly		bool	`json:"memberOnly"`	
	Url            	string 	`json:"url"`
}

func FromDomain(domain videos.Domain) Videos {
	return Videos{
		ID             	: domain.ID,
		Title          	: domain.Title,
		Classification 	: domain.ClassificationName,
		AdminID        	: domain.AdminID,
		MemberOnly		: domain.MemberOnly,
		Url            	: domain.Url,
	}
}

func FromDomainArray(domain []videos.Domain) []Videos {
	res := []Videos{}
	for _, val := range domain {
		res = append(res, FromDomain(val))
	}
	return res
}