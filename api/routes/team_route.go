package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"team-finder/api/controller"
	"team-finder/boot"
	"team-finder/domain"
	"team-finder/repository"
	"team-finder/usecase"
	"time"
)

func NewTeamRouter(env *boot.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	tr := repository.NewTeamRepository(db, domain.TableTeam)
	sc := controller.TeamController{
		TeamUsecase: usecase.NewTeamUsecase(tr, timeout),
		Env:         env,
	}
	group.GET("/teams/user/:id", sc.GetTeamsByUser)
	group.GET("/team/:id", sc.GetTeamById)
	group.GET("/teams/event/:id", sc.GetTeamsByEvent)
	group.GET("/teams", sc.GetAllTeams)

	group.POST("/teams", sc.RegTeam)
	group.PUT("/teams/:team_id/add-user/:user_id", sc.AddUserToTeam)
	group.POST("/teams/filter", sc.FilterTeams)
}
