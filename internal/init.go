package internal

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func Init(r *gin.Engine) {
	r.GET("/users/:id", getUserById)
	r.GET("/tags/user/:id", getTagsByUser)
	r.GET("/teams/user/:id", getTeamsByUser)
	r.GET("/team/:id", getTeamById)
	r.GET("/tags/team/:id", getTagsByTeam)
	r.GET("/tags", getAllTags)

	r.GET("/users/team/:id", getUsersByTeam)
	r.GET("/event/:id", getEventById)
	r.GET("/teams/event/:id", getTeamsByEvent)
	r.GET("/tags/event/:id", getTagsByEvent)
	r.GET("/events", getAllEvents)
	r.GET("/teams", getAllTeams)
	r.GET("/global-tags", getAllGlobalTags)
	r.GET("/tags/global-tag/:id", getTagsByGlobalTag)
	r.GET("/users", getAllUsers)

	r.POST("/users", regUser)
	r.POST("/events", regEvent)
	r.POST("/teams", regTeam)
}
