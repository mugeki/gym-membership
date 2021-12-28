package request

type TransactionClass struct {
	ID      uint
	UserID  uint   `json:"userID" valid:"required"`
	AdminID uint   `json:"adminID"`
	Status  string `json:"status"`
	Nominal int    `json:"nominal"`
	ClassID int    `json:"classID" valid:"required"`
}
