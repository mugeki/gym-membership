package response

import "time"

type MembershipTransaction struct {
	ID                  uint      `json:"id"`
	UserID              uint      `json:"user_id"`
	AdminID             uint      `json:"admin_id"`
	Status              string    `json:"status"`
	Nominal             int       `json:"nominal"`
	MembershipProductID int       `json:"membership_product_id"`
	CreatedAt           time.Time `json:"created_at"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
