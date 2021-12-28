package members

import (

	"gorm.io/gorm"
)

type members struct {
	gorm.Model
	ID 				uint
	// UUID      uuid.UUID
	Name			string		
	Url_image		string		
	Price			int			
	Period_time		time.time	
}
