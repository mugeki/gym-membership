package members

// import (
// 	"time"
// )

type Domain struct {
	ID 				uint
	// UUID      uuid.UUID
	Name			string
	Url_image		string
	Price			int
	Period_time		string
}

type Usecase interface {
	Insert(membersData *Domain) (string, error)
	GetByUserID(idMembers int) (string, error)
}

type Repository interface {
	Insert(membersData *Domain) (Domain, error)
	GetByUserID(idMembers int) (string, error)
	UpdateStatus(idMembers int) (string, error)
}
