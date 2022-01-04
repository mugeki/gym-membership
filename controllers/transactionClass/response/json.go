package response

type TransactionClass struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	AdminID uint   `json:"admin_id"`
	Status  string `json:"status"`
	Nominal int    `json:"nominal"`
	ClassID int    `json:"class_id"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}

type Class struct {
	ID uint
	// UUID      uuid.UUID
	Name         string `json:"name"`
	UrlImage     string `json:"url_image"`
	TrainerName  string `json:"trainer_name"`
	TrainerImage string `json:"trainer_image"`
	Description  string `json:"description"`
	IsOnline     bool   `json:"is_online"`
	StartDate    string `json:"start_date"`
	StartEnd     string `json:"start_end"`
	Location     string `json:"location"`
}
