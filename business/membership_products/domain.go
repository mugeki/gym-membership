package membership_products

type Domain struct {
	ID         uint
	Name       string
	UrlImage   string
	Price      int
	PeriodTime int
}

type Usecase interface {
	Insert(membershipProductsData *Domain) (string, error)
	GetByID(idMembers uint) (Domain, error)
}

type Repository interface {
	Insert(membershipProductsData *Domain) (Domain, error)
	GetByID(idMembers uint) (Domain, error)
}
