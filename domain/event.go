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

type EventRegRequest struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Date        string `form:"date" binding:"required"`
	Online      string `form:"online" binding:"required"`
	MainTheme   string `form:"main_theme" binding:"required"`
	Url         string `form:"url" binding:"required"`
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
}

type EventUsecase interface {
	GetAll() ([]Event, error)
	GetEventById(eventId int) (Event, error)
	RegEvent(request EventRegRequest, creatorId int) (int, error)
}
