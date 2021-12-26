package members

import (

	"gorm.io/gorm"
)

type members struct {
	gorm.Model
	ID 				uint
	// UUID      uuid.UUID
	UserID 			uint
	ExpiredDate		string
}
