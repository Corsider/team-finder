package repository

import (
	"database/sql"
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
	//TODO implement me
	panic("implement me")
}

func (t *teamRepository) GetByTeamId(id int) (domain.Team, error) {
	//TODO implement me
	panic("implement me")
}

func (t *teamRepository) GetByUserId(id int) ([]domain.Team, error) {
	//TODO implement me
	panic("implement me")
}

func (t *teamRepository) GetByEventId(id int) ([]domain.Team, error) {
	//TODO implement me
	panic("implement me")
}
