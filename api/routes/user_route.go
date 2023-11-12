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

func NewUserLoginRouter(env *boot.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.TableUser)
	sc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/users/login", sc.Login)
}
