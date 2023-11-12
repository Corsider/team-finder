package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"team-finder/api/routes"
	"team-finder/boot"
	"time"
)

func main() {
	app := boot.App()
	env := app.Env
	db := app.DB
	defer app.Close()
	timeout := time.Duration(env.Timeout) * time.Second

	r := gin.Default()

	routes.Init(env, timeout, db, r)

	err := r.Run(env.ServerAddress)
	if err != nil {
		log.Fatal("Error while starting gin engine", err)
	}
}
