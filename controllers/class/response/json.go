package response

type Class struct {
	ID uint
	// UUID      uuid.UUID
	Name            string `json:"name"`
	UrlImage        string `json:"urlImage"`
	Price           int    `json:"price"`
	Kuota           int    `json:"kuota"`
	TrainerName     string `json:"trainerName"`
	TrainerImage    string `json:"trainerImage"`
	Description     string `json:"description"`
	AvailableStatus bool   `json:"availableStatus"`
	IsOnline        bool   `json:"isOnline"`
	Date            string `json:"date"`
	Location        string `json:"location"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
