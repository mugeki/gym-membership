package admins

import "time"

type Domain struct {
	ID           uint
	Username     string
	Password     string
	Email        string
	FullName     string
	Gender       string
	Telephone    string
	Address      string
	UrlImage     string
	Token        string
	IsSuperAdmin bool
	CreatedAt    time.Time
}

type Usecase interface {
	Register(adminData *Domain) (Domain, error)
	Login(username, password string) (Domain, error)
	Update(id uint, adminData *Domain) (Domain, error)
	GetAll(id uint, title string, page int) ([]Domain, int, int, int64, error)
}

type Repository interface {
	Register(adminData *Domain) (Domain, error)
	GetByUsername(username string) (Domain, error)
	Update(id uint, adminData *Domain) (Domain, error)
	GetAll(id uint, title string, offset, limit int) ([]Domain, int64, error)
}
