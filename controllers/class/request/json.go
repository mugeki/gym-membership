package request

type Class struct {
	ID uint
	// UUID      uuid.UUID
	Name            string `json:"name" valid:"required"`
	UrlImage        string `json:"url_image" valid:"required"`
	Price           int    `json:"price" valid:"required"`
	Kuota           int    `json:"kuota" valid:"required"`
	TrainerId       int    `json:"trainer_id" valid:"required"`
	Description     string `json:"description" valid:"required"`
	AvailableStatus bool   `json:"available_status"`
	IsOnline        bool   `json:"is_online"`
	Date            string `json:"date" valid:"required"`
	Location        string `json:"location" valid:"required"`
}

type ClassUpdate struct {
	ID              uint
	Name            string `json:"name"`
	UrlImage        string `json:"url_image"`
	Price           int    `json:"price"`
	Kuota           int    `json:"kuota"`
	TrainerId       int    `json:"trainer_id"`
	Description     string `json:"description"`
	AvailableStatus bool   `json:"available_status"`
	IsOnline        bool   `json:"is_online"`
	Date            string `json:"date"`
	Location        string `json:"location"`
}
