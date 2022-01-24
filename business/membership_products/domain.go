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
	GetAll(page int) ([]Domain, int, int, int64, error)
	GetByID(id uint) (Domain, error)
	UpdateByID(id uint, newData *Domain) (Domain, error)
	DeleteByID(id uint) error
}

type Repository interface {
	Insert(newData *Domain) (Domain, error)
	GetAll(offset, limit int) ([]Domain, int64, error)
	GetByID(id uint) (Domain, error)
	UpdateByID(id uint, newData *Domain) (Domain, error)
	DeleteByID(id uint) error
}
