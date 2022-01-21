package response

import "time"

type Videos struct {
	ID                 uint      `json:"id"`
	Title              string    `json:"title"`
	ClassificationID   int		 `json:"classification_id"`
	ClassificationName string    `json:"classification"`
	AdminID            uint      `json:"admin_id"`
	MemberOnly         bool      `json:"member_only"`
	Url                string    `json:"url"`
	CreatedAt          time.Time `json:"created_at"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}