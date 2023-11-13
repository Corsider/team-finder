package postgres

import (
	"database/sql"
	"log"
	"team-finder/internal/utils"
)

type Database interface {
	Reference() *sql.DB
	SelectAllFromX(X string) (*sql.Rows, error)
	SelectAllFromXWhereYeqZ(X, Y, Z string) (*sql.Rows, error)
	SelectXFromYWhereZeqN(X, Y, Z, N string) (*sql.Rows, error)
	SelectXFromYWhereZeqNorMeqC(X, Y, Z, N, M, C string) (*sql.Rows, error)
	DeleteFromXWhereYeqZ(X, Y, Z string) error
	InsertIntoXYValuesZ(X, Y, Z string) error
	InsertIntoXYValuesZReturningN(X, Y, Z string) (interface{}, error)
}

type PostgresDB struct {
	DB *sql.DB
}

func (db *PostgresDB) Reference() *sql.DB {
	return db.DB
}

func NewConnection(connection string) Database {
	database, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	return &PostgresDB{DB: database}
}

func (db *PostgresDB) SelectAllFromX(X string) (*sql.Rows, error) {
	query := "select * from $1"
	return db.DB.Query(query, X)
}

func (db *PostgresDB) SelectAllFromXWhereYeqZ(X, Y, Z string) (*sql.Rows, error) {
	query := "select * from $1 where $2=$3"
	return db.DB.Query(query, X, Y, Z)
}

func (db *PostgresDB) SelectXFromYWhereZeqN(X, Y, Z, N string) (*sql.Rows, error) {
	query := "select $1 from $2 where $3=$4"
	return db.DB.Query(query, X, Y, Z, N)
}

func (db *PostgresDB) SelectXFromYWhereZeqNorMeqC(X, Y, Z, N, M, C string) (*sql.Rows, error) {
	query := "select $1 from $2 where $3=$4 or $5=$6"
	return db.DB.Query(query, X, Y, Z, N, M, C)
}

func (db *PostgresDB) DeleteFromXWhereYeqZ(X, Y, Z string) error {
	query := "delete from $1 where $2=$3"
	return utils.Second(db.DB.Exec(query, X, Y, Z))
}

func (db *PostgresDB) InsertIntoXYValuesZ(X, Y, Z string) error {
	query := "insert into $1($2) values ($3)"
	return utils.Second(db.DB.Exec(query, X, Y, Z))
}

func (db *PostgresDB) InsertIntoXYValuesZReturningN(X, Y, Z string) (interface{}, error) {
	query := "insert into $1($2) values ($3) returning $4"
	var returned interface{}
	err := db.DB.QueryRow(query, X, Y, Z).Scan(&returned)
	return returned, err
}
