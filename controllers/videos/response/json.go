package response

type Videos struct {
	ID                 uint   `json:"id"`
	Title              string `json:"title"`
	ClassificationName string `json:"classification"`
	AdminID            uint   `json:"admin_id"`
	MemberOnly         bool   `json:"member_only"`
	Url                string `json:"url"`
}