package domain

const (
	TableEvent = "event"
)

type Event struct {
	EventID     int    `db:"event_id" json:"event_id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Date        string `db:"date" json:"date"`
	Online      bool   `db:"online" json:"online"`
	MainTheme   string `db:"main_theme" json:"main_theme"`
	Url         string `db:"url" json:"url"`
	CreatorID   int    `db:"creator_id" json:"creator_id"`
}

type EventsResponse struct {
	Events []Event `json:"events"`
}

type EventResponse struct {
	Event Event `json:"event"`
}

type EventRepository interface {
	GetAll() ([]Event, error)
	GetEventById(eventId int) (Event, error)
}

type EventUsecase interface {
	GetAll() ([]Event, error)
	GetEventById(eventId int) (Event, error)
}
