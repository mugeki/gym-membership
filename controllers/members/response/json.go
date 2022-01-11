package response

import "time"

type Member struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	ExpireDate time.Time `json:"expire_date"`
	CreatedAt  time.Time `json:"created_at"`
}