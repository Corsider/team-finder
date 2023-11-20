package repository

import (
	"database/sql"
	"strconv"
	"team-finder/domain"
	"team-finder/internal/utils"
	"team-finder/postgres"
)

type teamRepository struct {
	database postgres.Database //*sql.DB
	table    string
}

func NewTeamRepository(db *sql.DB, table string) domain.TeamRepository {
	return &teamRepository{
		database: &postgres.PostgresDB{DB: db},
		table:    table,
	}
}

func (t *teamRepository) GetAll() ([]domain.Team, error) {
	rows, err := t.database.SelectAllFromX(t.table)
	if err != nil {
		return nil, err
	}
	teams := []domain.Team{}
	for rows.Next() {
		var team domain.Team
		rows.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
		teams = append(teams, team)
	}
	return teams, nil
}

func (t *teamRepository) GetByTeamId(id int) (domain.Team, error) {
	row := t.database.Select1FromXWhereYeqZ(t.table, "team_id", strconv.Itoa(id))
	var team domain.Team
	err := row.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
	if err != nil {
		return domain.Team{}, err
	}
	return team, nil
}

func (t *teamRepository) GetByUserId(id int) ([]domain.Team, error) {
	rows, err := t.database.SelectAllFromXWhereYeqZ("user_team", "user_id", strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	teams := []domain.Team{}
	var userTeam domain.UserTeam
	for rows.Next() {
		var team domain.Team
		rows.Scan(&userTeam.TeamId, &userTeam.UserId, &userTeam.Role, &userTeam.DateOfEntry, &userTeam.Hidden)
		row := t.database.Select1FromXWhereYeqZ("team", "team_id", strconv.Itoa(userTeam.TeamId))
		row.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
		teams = append(teams, team)
	}
	return teams, nil
}

func (t *teamRepository) GetByEventId(id int) ([]domain.Team, error) {
	rows, err := t.database.SelectAllFromXWhereYeqZ("team_event", "event_id", strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	teams := []domain.Team{}
	var teamEvent domain.TeamEvent
	for rows.Next() {
		var team domain.Team
		rows.Scan(&teamEvent.EventId, &teamEvent.TeamId, &teamEvent.RegTime)
		row := t.database.Select1FromXWhereYeqZ("team", "team_id", strconv.Itoa(teamEvent.TeamId))
		row.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
		teams = append(teams, team)
	}
	return teams, nil
}

func (t *teamRepository) RegTeam(request domain.TeamsRegRequest) (int, error) {
	//teamId, err := t.database.InsertIntoXYValuesZReturningN(t.table, "name, description, rules, place, reg_date, rate",
	//	request.Name+", "+request.Description+", "+request.Rules+", "+request.Place+", CURRENT_TIMESTAMP, 5", "team_id")
	teamId, err := t.database.InsertParametrizedIntoXYValuesZReturningN(t.table, "name, description, rules, place, reg_date, rate",
		"$1, $2, $3, $4, CURRENT_TIMESTAMP, 5", "team_id", request.Name, request.Description, request.Rules, request.Place)
	if err != nil {
		return 0, err
	}
	return utils.First(strconv.Atoi(strconv.FormatInt(teamId.(int64), 10))), nil
}

func (t *teamRepository) AddUserToTeam(userId int, teamId int) error {
	//err := t.database.InsertIntoXYValuesZ("user_team", "team_id, user_id, role, date_of_entry, hidden",
	//	strconv.Itoa(teamId)+", "+strconv.Itoa(userId)+", "+"Creator, CURRENT_TIMESTAMP, false")
	err := t.database.InsertParametrizedIntoXYValuesZ("user_team", "team_id, user_id, role, date_of_entry, hidden",
		"$1, $2, 'Creator', CURRENT_TIMESTAMP, false", teamId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (t *teamRepository) DeleteTeamById(teamId int) error {
	return t.database.DeleteFromXWhereYeqZ(t.table, "team_id", strconv.Itoa(teamId))
}

func (t *teamRepository) FilterTeams(onlyUser bool, tags []int, myTeam int, sortBy string, asc bool, from, to int) ([]domain.Team, error) {
	var order string
	switch sortBy {
	case "count":
		order = "(select count(*) from user_team where team_id = team.team_id)"
	case "activity":
		order = "(select count(*) from team_event where team_id = team.team_id)"
	case "name":
		order = "name"
	}
	if onlyUser {
		rows, err := t.database.SelectAllFromXWhereYinZandNinMandGinHOrderByO(t.table, "team_id", "select team_id from user_team where user_id=$1",
			"team_id", "select team_id from team_tags where tag_id in $2", "select count(*) from user_team where team_id=team.team_id",
			"$3 and $4", order, myTeam, tags, from, to, asc)
		if err != nil {
			return nil, err
		}
		teams := []domain.Team{}
		for rows.Next() {
			var team domain.Team
			rows.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
			teams = append(teams, team)
		}
		return teams, nil
	} else {
		rows, err := t.database.SelectAllFromXNinMandGinHOrderByO(t.table, "team_id", "select team_id from team_tags where tag_id in $1",
			"select count(*) from user_team where team_id=team.team_id", "$2 and $3", order, myTeam, tags, from, to, asc)
		if err != nil {
			return nil, err
		}
		teams := []domain.Team{}
		for rows.Next() {
			var team domain.Team
			rows.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
			teams = append(teams, team)
		}
		return teams, nil
	}
}

func (t *teamRepository) FilterTeamUser(order string, tags []int, myTeam int, asc bool, from, to int) ([]domain.Team, error) {
	rows, err := t.database.SelectAllFromXWhereYinZandNinMandGinHOrderByO(t.table, "team_id", "select team_id from user_team where user_id=$1",
		"team_id", "select team_id from team_tags where tag_id in $2", "select count(*) from user_team where team_id=team.team_id",
		"$3 and $4", order, myTeam, tags, from, to, asc)
	if err != nil {
		return nil, err
	}
	teams := []domain.Team{}
	for rows.Next() {
		var team domain.Team
		rows.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
		teams = append(teams, team)
	}
	return teams, nil
}

func (t *teamRepository) FilterTeamNoUser(order string, tags []int, myTeam int, asc bool, from, to int) ([]domain.Team, error) {
	rows, err := t.database.SelectAllFromXNinMandGinHOrderByO(t.table, "team_id", "select team_id from team_tags where tag_id in $1",
		"select count(*) from user_team where team_id=team.team_id", "$2 and $3", order, myTeam, tags, from, to, asc)
	if err != nil {
		return nil, err
	}
	teams := []domain.Team{}
	for rows.Next() {
		var team domain.Team
		rows.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
		teams = append(teams, team)
	}
	return teams, nil
}
