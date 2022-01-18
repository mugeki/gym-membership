package payment_accounts

type Domain struct {
	ID        uint
	Name      string
	NoCard    string
	OwnerName string
	Desc      string
}

type Usecase interface {
	Insert(account *Domain) (Domain, error)
	GetAll() ([]Domain, error)
}

type Repository interface {
	Insert(account *Domain) (Domain, error)
	GetAll() ([]Domain, error)
}
