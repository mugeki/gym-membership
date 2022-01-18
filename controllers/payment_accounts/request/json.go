package request

type PaymentAccount struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	NoCard    string `json:"no_card"`
	OwnerName string `json:"owner_name"`
	Desc      string `json:"desc"`
}
