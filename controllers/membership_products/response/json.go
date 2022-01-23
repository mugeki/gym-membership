package response

type MembershipProducts struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	UrlImage   string `json:"url_image"`
	Price      int    `json:"price"`
	PeriodTime int    `json:"period_time"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}