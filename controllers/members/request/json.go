package request

type Members struct {
	ID uint
	// UUID      uuid.UUID
	Name			string		`json:"name" valid:"required"`
	Url_image		string		`json:"url_image" valid:"required"`
	Price			int			`json:"price" valid:"required"`
	Period_time		string	`json:"period_time" valid:"required"`
}
	