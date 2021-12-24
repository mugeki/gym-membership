package response

type Articles struct {
	ID                 uint   `json:"id"`
	Title              string `json:"title"`
	ClassificationName string `json:"classification"`
	AdminID            uint   `json:"admin_id"`
	MemberOnly         bool   `json:"member_only"`
	Url                string `json:"url"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
