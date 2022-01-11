package response

type MembershipProducts struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	UrlImage   string `json:"url_image"`
	Price      int    `json:"price"`
	PeriodTime int    `json:"period_time"`
}