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

func NewEventRouter(env *boot.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	tr := repository.NewEventRepository(db, domain.TableEvent)
	sc := controller.EventController{
		EventUsecase: usecase.NewEventUsecase(tr, timeout),
		Env:          env,
	}
	group.GET("/event/:id", sc.GetEventById)
	group.GET("/events", sc.GetAll)
	group.POST("/events", sc.RegEvent)

	group.PUT("/events/:event_id/add-team/:team_id", sc.AddTeamToEvent)
	group.DELETE("/events/:id", sc.DeleteEvent)
}
