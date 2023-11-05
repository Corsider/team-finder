package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"team-finder/internal"
	"team-finder/services"
)

var DB *sql.DB

func main() {
	log.Println("Connecting to postgres...")
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}
	DB = services.Connect(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))
	defer DB.Close()

	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("CONNECTED")
		internal.DB = DB
	}

	r := gin.Default()
	internal.Init(r)
	err = r.Run(os.Getenv("HOST"))
	if err != nil {
		log.Fatal(err)
	}
}
