package request

type MembershipTransaction struct {
	UserID              uint   `json:"user_id" valid:"required"`
	AdminID             uint   `json:"admin_id"`
	Status              string `json:"status"`
	Nominal             int    `json:"nominal"`
	MembershipProductID uint   `json:"membership_product_id" valid:"required"`
}
