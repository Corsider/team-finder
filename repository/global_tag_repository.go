package repository

import (
	"database/sql"
	"team-finder/domain"
	"team-finder/postgres"
)

type globalTagRepository struct {
	database postgres.Database //*sql.DB
	table    string
}

func NewGlobalTagRepository(db *sql.DB, table string) domain.GlobalTagRepository {
	return &globalTagRepository{
		database: &postgres.PostgresDB{DB: db},
		table:    table,
	}
}

func (gt *globalTagRepository) GetAll() ([]domain.GlobalTag, error) {
	rows, err := gt.database.SelectAllFromX(gt.table)
	if err != nil {
		return nil, err
	}
	gtags := []domain.GlobalTag{}
	for rows.Next() {
		var gtag domain.GlobalTag
		rows.Scan(&gtag.GlobalTagID, &gtag.Category)
		gtags = append(gtags, gtag)
	}
	return gtags, err
}
