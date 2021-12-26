package request

type Class struct {
	ID uint
	// UUID      uuid.UUID
	Name            string `json:"name" valid:"required"`
	UrlImage        string `json:"urlImage" valid:"required"`
	Price           int    `json:"price" valid:"required"`
	Kuota           int    `json:"kuota" valid:"required"`
	TrainerId       int    `json:"trainerId" valid:"required"`
	Description     string `json:"description" valid:"required"`
	AvailableStatus bool   `json:"availableStatus"`
	IsOnline        bool   `json:"isOnline"`
	Date            string `json:"date" valid:"required"`
	Location        string `json:"location" valid:"required"`
}
