package request

type TransactionClass struct {
	ID       uint
	UserID   uint   `json:"userID" valid:"required"`
	AdminID  uint   `json:"adminID"`
	Status   string `json:"status"`
	Nominal  string `json:"nominal" valid:"required"`
	ClassID  int    `json:"classID" valid:"required"`
	Location string `json:"location" valid:"required"`
}
