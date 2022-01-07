package members

import "time"

type Domain struct {
	ID uint
	UserID	uint
	ExpireDate time.Time
	CreatedAt time.Time
}

type Usecase interface{
	GetByUserID(userID uint) (Domain, error)
}

type Repository interface{
	Insert(data *Domain) error
	GetByUserID(userID uint) (Domain, error)
}