package internal

import "github.com/lib/pq"

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

type GlobalTag struct {
	GlobalTagID int    `db:"globaltag_id" json:"globaltag_id"`
	Category    string `db:"category" json:"category"`
}

type Tag struct {
	TagID       int    `db:"tag_id" json:"tag_id"`
	Activity    string `db:"activity" json:"activity"`
	GlobalTagID int    `db:"globaltag_id" json:"globaltag_id"`
}

type Team struct {
	TeamID      int     `db:"team_id" json:"team_id"`
	Name        string  `db:"name" json:"name"`
	Rate        float32 `db:"rate" json:"rate"`
	Description string  `db:"description" json:"description"`
	Rules       string  `db:"rules" json:"rules"`
	RegDate     string  `db:"reg_date" json:"reg_date"`
	Place       string  `db:"place" json:"place"`
}

type User struct {
	UserId      int     `db:"user_id" json:"user_id"`
	Name        string  `db:"name" json:"name"`
	Nickname    string  `db:"nickname" json:"nickname"`
	Rate        float64 `db:"rate" json:"rate"`
	Description string  `db:"description" json:"description"`
	Login       string  `db:"login" json:"login"`
	Password    string  `db:"password" json:"password"`
}

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

type TagArray struct {
	Tags pq.Int32Array `json:"tags"`
}
