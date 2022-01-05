package request

type MembershipProducts struct {
	Name       string `json:"name" valid:"required"`
	UrlImage   string `json:"url_image" valid:"required"`
	Price      int    `json:"price" valid:"required"`
	PeriodTime int    `json:"period_time" valid:"required"`
}