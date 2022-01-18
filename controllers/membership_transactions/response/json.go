package response

import "time"

type MembershipTransaction struct {
	ID                  uint           `json:"id"`
	UserID              uint           `json:"user_id"`
  UserName			      string	  `json:"user_name"`
	AdminID             uint           `json:"admin_id"`
	Status              string         `json:"status"`
	Nominal             int            `json:"nominal"`
	MembershipProductID int            `json:"membership_product_id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	UrlImageOfReceipt   string         `json:"url_image_of_receipt"`
	Payment             PaymentAccount `json:"payment"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}

type PaymentAccount struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	NoCard    string `json:"no_card"`
	OwnerName string `json:"owner_name"`
	Desc      string `json:"desc"`
}
