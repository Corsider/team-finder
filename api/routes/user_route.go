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

func NewUserLogRouter(env *boot.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.TableUser)
	sc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		Env:         env,
	}
	group.POST("/users/login", sc.Login)
	group.POST("/users/signup", sc.Signup) // todo
}

func NewUserRouter(env *boot.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.TableUser)
	sc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		Env:         env,
	}
	group.GET("/users/:id", sc.GetById)
	group.GET("/users/team/:id", sc.GetUsersByTeamId)
	group.GET("/users", sc.GetAll)
}
