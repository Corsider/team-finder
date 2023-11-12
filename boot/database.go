package boot

import (
	"database/sql"
	"fmt"
	"log"
)

func NewPostgresDB(env *Env) *sql.DB {
	psinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName)
	database, err := sql.Open("postgres", psinfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Postgres")
	return database
}

func ClosePostgresConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to Postgres closed")
}
