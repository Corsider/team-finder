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

func NewTagRouter(env *boot.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	tr := repository.NewTagRepository(db, domain.TableTag)
	sc := controller.TagController{
		TagUsecase: usecase.NewTagUsecase(tr, timeout),
		Env:        env,
	}
	group.GET("/tags", sc.GetAll)
	group.GET("/tags/user/:id", sc.GetByUserId)
	group.GET("/tags/team/:id", sc.GetByTeamId)
	group.GET("/tags/event/:id", sc.GetByEventId)
	group.GET("/tags/global-tag/:id", sc.GetByGlobalTagId)

	group.POST("/tags/user", sc.PostTagsToUser)
	group.POST("/tags/team", sc.PostTagsToTeam)
	group.POST("/tags/event", sc.PostTagsToEvent)

	group.DELETE("/tags/user/:id", sc.DeleteTagsFromUser)
	group.DELETE("/tags/team/:id", sc.DeleteTagsFromTeam)
	group.DELETE("/tags/event/:id", sc.DeleteTagsFromEvent)
}
