package services

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect(user, password, dbMame string) *sql.DB {
	psinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "postgres", "5432", user, password, dbMame)
	database, err := sql.Open("postgres", psinfo)
	if err != nil {
		log.Fatal(err)
	}
	return database
}
