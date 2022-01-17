package request

type ClassTransaction struct {
	ID        uint
	UserID    uint   `json:"user_id" valid:"required"`
	AdminID   uint   `json:"admin_id"`
	Status    string `json:"status"`
	Nominal   int    `json:"nominal"`
	ClassID   int    `json:"class_id" valid:"required"`
	PaymentID int    `json:"payment_id" valid:"required"`
}
