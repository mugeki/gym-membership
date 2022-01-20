package membership_products

type Domain struct {
	ID         uint
	Name       string
	UrlImage   string
	Price      int
	PeriodTime int
}

type Usecase interface {
	Insert(newData *Domain) (Domain, error)
	GetAll() ([]Domain, error)
	GetByID(id uint) (Domain, error)
	UpdateByID(id uint, newData *Domain) (Domain, error)
	DeleteByID(id uint) error
}

type Repository interface {
	Insert(newData *Domain) (Domain, error)
	GetAll() ([]Domain, error)
	GetByID(id uint) (Domain, error)
	UpdateByID(id uint, newData *Domain) (Domain, error)
	DeleteByID(id uint) error
}
