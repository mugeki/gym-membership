package request

type Articles struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	ClassificationID uint   `json:"classification_id"`
	AdminID          uint   `json:"admin_id"`
	MemberOnly       bool   `json:"member_only"`
	UrlImage         string `json:"url_image"`
	Text             string `json:"text"`
}
