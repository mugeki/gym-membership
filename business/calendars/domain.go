package calendars

type Domain struct {
	NameEvent string
	StartDate string
	EndDate   string
	Location  string
}

type Usecase interface {
	Insert(EventData *Domain) (string, error)
}

type Repository interface {
	Insert(EventData *Domain) (Domain, error)
}
