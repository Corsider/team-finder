package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"team-finder/api/middleware"
	"team-finder/boot"
	"time"
)

func Init(env *boot.Env, timeout time.Duration, DB *sql.DB, r *gin.Engine) {
	public := r.Group("")
	NewUserLoginRouter(env, timeout, DB, public)
	//
	protected := r.Group("")
	protected.Use(middleware.JWT(env.TokenSecret))
	//
}
