package repository

import (
	"database/sql"
	"team-finder/domain"
	"team-finder/postgres"
)

type eventRepository struct {
	database postgres.Database //*sql.DB
	table    string
}

func NewEventRepository(db *sql.DB, table string) domain.EventRepository {
	return &eventRepository{
		database: &postgres.PostgresDB{DB: db},
		table:    table,
	}
}

func (er *eventRepository) GetAll() ([]domain.Event, error) {
	// todo
	panic("")
}

func (er *eventRepository) GetEventById(eventId int) (domain.Event, error) {
	// todo
	panic("")
}
