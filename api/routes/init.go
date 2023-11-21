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
	NewUserLogRouter(env, timeout, DB, public)
	//
	protected := r.Group("")
	protected.Use(middleware.JWT(env.TokenSecret))
	NewUserRouter(env, timeout, DB, protected)
	NewTagRouter(env, timeout, DB, protected)
	NewTeamRouter(env, timeout, DB, protected)
	NewEventRouter(env, timeout, DB, protected)
	NewGlobalTagRouter(env, timeout, DB, protected)
}
