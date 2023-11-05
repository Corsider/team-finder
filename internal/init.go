package internal

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func Init(r *gin.Engine) {
	r.GET("/setup", setup)
	r.GET("/users/:id", getUserById)
	r.GET("/tags/user/:id", getTagsByUser)
	r.GET("/teams/user/:id", getTeamsByUser)
	r.GET("/team/:id", getTeamById)
	r.GET("/tags/team/:id", getTagsByTeam)
	r.GET("/tags", getAllTags)
}
