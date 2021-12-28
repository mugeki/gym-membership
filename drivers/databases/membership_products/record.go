package membership_products

import (

	"gorm.io/gorm"
)

type MembershipProducts struct {
	gorm.Model
	ID 				uint
	// UUID      uuid.UUID
	Name			string		
	UrlImage		string		
	Price			int			
	PeriodTime		int	
}
