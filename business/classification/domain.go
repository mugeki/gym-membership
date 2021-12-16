package classification

type Domain struct {
	ID   uint
	Name string
}

type Usecase interface {
	GetClassificationID(name string) (uint, error)
}

type Repository interface {
	GetClassificationID(name string) (uint, error)
}
