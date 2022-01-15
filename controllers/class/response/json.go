package response

type Class struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	UrlImage        string `json:"url_image"`
	Price           int    `json:"price"`
	Kuota           int    `json:"kuota"`
<<<<<<< HEAD
	Participant     int    `json:"participant"`
=======
>>>>>>> a298db310b1b5d2007ff3b9fab36b7cebe050a27
	TrainerName     string `json:"trainer_name"`
	TrainerImage    string `json:"trainer_image"`
	Description     string `json:"description"`
	AvailableStatus bool   `json:"available_status"`
	IsOnline        bool   `json:"is_online"`
	Date            string `json:"date"`
	Location        string `json:"location"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
