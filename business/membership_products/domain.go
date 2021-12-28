package membership_products
// import (
// 	"time"
// )

type Domain struct {
	ID 				uint
	// UUID      uuid.UUID
	Name			string		
	UrlImage		string		
	Price			int			
	PeriodTime		int	
}

type Usecase interface {
	Insert(membershipProductsData *Domain) (string, error)
	GetByUserID(idMembers uint) (string, error)
}

type Repository interface {
	Insert(membershipProductsData *Domain) (Domain, error)
	GetByUserID(idMembers uint) (string, error)
}
