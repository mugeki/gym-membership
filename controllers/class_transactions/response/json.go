package response

import "time"

type ClassTransaction struct {
	ID                uint           `json:"id"`
	UserID            uint           `json:"user_id"`
  	UserName	      string	  	 `json:"user_name"`
	AdminID           uint           `json:"admin_id"`
	Status            string         `json:"status"`
	Nominal           int            `json:"nominal"`
	ClassID           int            `json:"class_id"`
	ClassName		  string		 `json:"class_name"`
	UrlImageOfReceipt string         `json:"url_image_of_receipt"`
	UpdatedAt         time.Time      `json:"updated_at"`
	CreatedAt         time.Time      `json:"created_at"`
	Payment           PaymentAccount `json:"payment"`
}

type ClassActive struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Date            string `json:"date"`
	Location        string `json:"location"`
	UrlImage		string `json:"url_image"`
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
