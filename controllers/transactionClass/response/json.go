package response

type TransactionClass struct {
	ID       uint
	UserID   uint   `json:"userID"`
	AdminID  uint   `json:"adminID"`
	Status   string `json:"status"`
	Nominal  string `json:"nominal"`
	ClassID  int    `json:"classID"`
	Location string `json:"location"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
