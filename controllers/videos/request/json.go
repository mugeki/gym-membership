package request

type Videos struct {
	Title              string `json:"title" valid:"required"`
	ClassificationName string `json:"classification" valid:"required"`
	MemberOnly         bool   `json:"member_only"`
	Url                string `json:"url" valid:"required, url"`
}