package response

type MembershipTransaction struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"userID"`
	AdminID uint   `json:"adminID"`
	Status  string `json:"status"`
	Nominal int    `json:"nominal"`
	ClassID int    `json:"classID"`
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
	UrlImage     string `json:"urlImage"`
	TrainerName  string `json:"trainerName"`
	TrainerImage string `json:"trainerImage"`
	Description  string `json:"description"`
	IsOnline     bool   `json:"isOnline"`
	StartDate    string `json:"startDate"`
	StartEnd     string `json:"startEnd"`
	Location     string `json:"location"`
}
