package classification

type Domain struct {
	ID   uint
	Name string
}

type Usecase interface {
	Insert(classification *Domain) (Domain, error)
	GetAll() ([]Domain, error)
}

type Repository interface {
	Insert(classification *Domain) (Domain, error)
	GetAll() ([]Domain, error)
}
