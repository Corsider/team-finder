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

func NewGlobalTagRouter(env *boot.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	tr := repository.NewGlobalTagRepository(db, domain.TableGlobalTag)
	sc := controller.GlobalTagController{
		GlobaTagUsecase: usecase.NewGlobalTagUsecase(tr, timeout),
		Env:             env,
	}
	group.GET("/global-tags", sc.GetAll)
}
