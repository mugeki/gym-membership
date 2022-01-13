package response

import "time"

type ClassTransaction struct {
	ID        	uint      `json:"id"`
	UserID    	uint      `json:"user_id"`
	AdminID   	uint      `json:"admin_id"`
	Status    	string    `json:"status"`
	Nominal   	int       `json:"nominal"`
	ProductName string	  `json:"product_name"`
	ClassID   	int       `json:"class_id"`
	CreatedAt 	time.Time `json:"created_at"`
}

type ClassActive struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Date            string `json:"date"`
	Location        string `json:"location"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
