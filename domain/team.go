package domain

const (
	TableTeam      = "team"
	TableTeamEvent = "team_event"
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
	RegTeam(request TeamsRegRequest) (int, error)
	AddUserToTeam(userId int, teamId int) error
	DeleteTeamById(teamId int) error
	FilterTeamUser(order string, tags []int, myTeam int, asc bool, from, to int) ([]Team, error)
	FilterTeamNoUser(order string, tags []int, myTeam int, asc bool, from, to int) ([]Team, error)
}

type TeamsAllResponse struct {
	Tags []Tag `json:"tags"`
}

type TeamsRegRequest struct {
	Name        string `form:"name" binding:"required"`
	Description string `from:"description" binding:"required"`
	Rules       string `form:"rules" binding:"required"`
	Place       string `from:"place" binding:"required"`
}

type TeamsRegResponse struct {
	TeamId string `json:"team_id"`
}

type TeamsFilterRequest struct {
	TagsId []int `json:"tags"`
}

type TeamUsecase interface {
	GetAll() ([]Team, error)
	GetByTeamId(id int) (Team, error)
	GetByUserId(id int) ([]Team, error)
	GetByEventId(id int) ([]Team, error)
	RegTeam(request TeamsRegRequest) (int, error)
	AddUserToTeam(userId int, teamId int) error
	DeleteTeamById(teamId int) error
	Filter(onlyUser bool, tags []int, myTeam int, sortBy string, asc bool, from, to int) ([]Team, error)
}
