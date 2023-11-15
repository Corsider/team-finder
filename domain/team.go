package domain

const (
	TableTeam = "team"
)

type Team struct {
	TeamID      int     `db:"team_id" json:"team_id"`
	Name        string  `db:"name" json:"name"`
	Rate        float32 `db:"rate" json:"rate"`
	Description string  `db:"description" json:"description"`
	Rules       string  `db:"rules" json:"rules"`
	RegDate     string  `db:"reg_date" json:"reg_date"`
	Place       string  `db:"place" json:"place"`
}

type TeamRepository interface {
	GetAll() ([]Team, error)
	GetByTeamId(id int) (Team, error)
	GetByUserId(id int) ([]Team, error)
	GetByEventId(id int) ([]Team, error)
}

type TeamsAllResponse struct {
	Tags []Tag `json:"tags"`
}

type TeamUsecase interface {
	GetAll() ([]Team, error)
	GetByTeamId(id int) (Team, error)
	GetByUserId(id int) ([]Team, error)
	GetByEventId(id int) ([]Team, error)
}
