package request

type MembershipTransaction struct {
	UserID              uint   `json:"user_id" valid:"required"`
	AdminID             uint   `json:"admin_id"`
	Status              string `json:"status"`
	Nominal             int    `json:"nominal"`
	MembershipProductID uint   `json:"membership_product_id" valid:"required"`
	UrlImageOfReceipt   string `json:"url_image_of_receipt"`
	PaymentID           uint   `json:"payment_id" valid:"required"`
}

type UpdateReceipt struct {
	UrlImageOfReceipt string `json:"url_image_of_receipt"`
}
