package request

import "time"

type Event struct {
	Kind     string    `json:"kind"`
	Etag     string    `json:"etag"`
	ID       string    `json:"id"`
	Status   string    `json:"status"`
	HTMLLink string    `json:"htmlLink"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Creator  struct {
		Email string `json:"email"`
	} `json:"creator"`
	Organizer struct {
		Email       string `json:"email"`
		DisplayName string `json:"displayName"`
		Self        bool   `json:"self"`
	} `json:"organizer"`
	Start struct {
		DateTime time.Time `json:"dateTime"`
		TimeZone string    `json:"timeZone"`
	} `json:"start"`
	End struct {
		DateTime time.Time `json:"dateTime"`
		TimeZone string    `json:"timeZone"`
	} `json:"end"`
	ICalUID   string `json:"iCalUID"`
	Sequence  int    `json:"sequence"`
	Attendees []struct {
		Email          string `json:"email"`
		ResponseStatus string `json:"responseStatus"`
	} `json:"attendees"`
	Reminders struct {
		UseDefault bool `json:"useDefault"`
	} `json:"reminders"`
	EventType string `json:"eventType"`
}
