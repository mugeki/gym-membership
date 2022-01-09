package response

type ClassTransaction struct {
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
