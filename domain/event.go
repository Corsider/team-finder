package domain

import "time"

const (
	TableEvent      = "events"
	TableEventsTags = "events_tags"
)

type Event struct {
	EventID     int       `db:"event_id" json:"event_id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Date        time.Time `db:"date" json:"date"`
	Online      bool      `db:"online" json:"online"`
	MainTheme   string    `db:"main_theme" json:"main_theme"`
	Url         string    `db:"url" json:"url"`
	CreatorID   int       `db:"creator_id" json:"creator_id"`
}

type EventRegRequest struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description"`
	Date        string `form:"date" binding:"required"`
	Online      bool   `form:"online"`
	MainTheme   string `json:"main_theme"`
	Url         string `form:"url"`
}

type EventRegResponse struct {
	EventId string `json:"event_id"`
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
	RegEvent(request EventRegRequest, creatorId int) (int, error)
	AddTeamToEvent(eventId, teamId int) error
	DeleteFromEvents(eventId int) error
	DeleteFromEventsTags(eventId int) error
	DeleteFromTeamEvent(eventId int) error
}

type EventUsecase interface {
	GetAll() ([]Event, error)
	GetEventById(eventId int) (Event, error)
	RegEvent(request EventRegRequest, creatorId int) (int, error)
	AddTeamToEvent(eventId, teamId int) error
	DeleteEvent(evenId int) error
}
