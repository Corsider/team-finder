package repository

import (
	"database/sql"
	"strconv"
	"team-finder/domain"
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
	for rows.Next() {
		var team domain.Team
		rows.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
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
	for rows.Next() {
		var team domain.Team
		rows.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
		teams = append(teams, team)
	}
	return teams, nil
}
