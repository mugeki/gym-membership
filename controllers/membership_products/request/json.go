package request

type MembershipProducts struct {
	ID uint
	// UUID      uuid.UUID
	Name			string		`json:"name" valid:"required"`
	UrlImage		string		`json:"url_image" valid:"required"`
	Price			int			`json:"price" valid:"required"`
	PeriodTime		int			`json:"period_time" valid:"required"`
}
	