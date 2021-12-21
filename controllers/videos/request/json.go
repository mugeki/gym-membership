package request

type Videos struct {
	Title            string `json:"title" valid:"required"`
	ClassificationID int    `json:"classification_id" valid:"required"`
	MemberOnly       bool   `json:"member_only"`
	AdminID          uint   `json:"admin_id"`
	Url              string `json:"url" valid:"required, url"`
}