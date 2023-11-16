package repository

import (
	"database/sql"
	"strconv"
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
	rows, err := er.database.SelectAllFromX(er.table)
	if err != nil {
		return nil, err
	}
	events := []domain.Event{}
	for rows.Next() {
		var event domain.Event
		rows.Scan(&event.EventID, &event.Name, &event.Description, &event.Date, &event.Online, &event.MainTheme, &event.Url, &event.CreatorID)
		events = append(events, event)
	}
	return events, nil
}

func (er *eventRepository) GetEventById(eventId int) (domain.Event, error) {
	rows := er.database.Select1FromXWhereYeqZ(er.table, "event_id", strconv.Itoa(eventId))
	var event domain.Event
	err := rows.Scan(&event.EventID, &event.Name, &event.Description, &event.Date, &event.Online, &event.MainTheme, &event.Url, &event.CreatorID)
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}
