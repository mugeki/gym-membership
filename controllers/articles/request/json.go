package request

type Articles struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	ClassificationID uint   `json:"classificationId"`
	AdminID          uint   `json:"adminId"`
	MemberOnly       bool   `json:"memberOnly"`
	UrlImage         string `json:"urlImage"`
	Text             string `json:"text"`
}
