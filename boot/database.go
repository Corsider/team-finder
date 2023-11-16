package boot

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"team-finder/postgres"
)

func NewPostgresDB(env *Env) *sql.DB {
	psinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.DBHost, strconv.Itoa(env.DBPort), env.DBUser, env.DBPass, env.DBName)
	database := postgres.NewConnection(psinfo)
	err := database.Reference().Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Postgres")
	return database.Reference()
}

func ClosePostgresConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to Postgres closed")
}
