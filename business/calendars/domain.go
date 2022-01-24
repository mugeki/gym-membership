package calendars

import "time"

type Event struct {
	Kind     string
	Etag     string
	ID       string
	Status   string
	HTMLLink string
	Created  time.Time
	Updated  time.Time
	Summary  string
	Creator  struct {
		Email string
	}
	Organizer struct {
		Email       string
		DisplayName string
		Self        bool
	}
	Start struct {
		DateTime time.Time
		TimeZone string
	}
	End struct {
		DateTime time.Time
		TimeZone string
	}
	ICalUID   string
	Sequence  int
	Attendees []struct {
		Email          string
		ResponseStatus string
	}
	Reminders struct {
		UseDefault bool
	}
	EventType string
}

type Token struct {
	AccessToken string
	TokenType   string
	Expiry      time.Time
}
type Usecase interface {
	CreateEvent(EventData *Event) (Event, error)
	AddGuest(eventId, emailGuest string) (Event, error)
	GetAll(code, state string) (Event, error)
}

type Repository interface {
	CreateEvent(EventData *Event) (Event, error)
	AddGuest(eventId, emailGuest string) (Event, error)
	GetAll(code, state string) (Event, error)
}
