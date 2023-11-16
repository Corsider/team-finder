package domain

type UsersTags struct {
	TagId  int `db:"tag_id" json:"tag_id"`
	UserId int `db:"user_id" json:"user_id"`
}

type EventTags struct {
	EventId int `json:"event_id" db:"event_id"`
	TagId   int `json:"tag_id" db:"tag_id"`
}

type TeamEvent struct {
	EventId int    `json:"event_id" db:"event_id"`
	TeamId  int    `json:"team_id" db:"team_id"`
	RegTime string `json:"reg_time" db:"reg_time"`
}

type TeamsTags struct {
	TagId  int `json:"tag_id" db:"tag_id"`
	TeamId int `json:"team_id" db:"team_id"`
}

type UserTeam struct {
	TeamId      int    `json:"team_id" db:"team_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	Role        string `json:"role" db:"role"`
	DateOfEntry string `json:"date_of_entry" db:"date_of_entry"`
	Hidden      bool   `json:"hidden" db:"hidden"`
}
