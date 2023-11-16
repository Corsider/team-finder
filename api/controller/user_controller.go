package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"team-finder/boot"
	"team-finder/domain"
	"team-finder/internal/utils"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *boot.Env
}

func (lc *UserController) Login(c *gin.Context) {
	var request domain.LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: err.Error()})
		return
	}
	user, err := lc.UserUsecase.GetUserByLogin(request.Login)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Error: "Not found"})
		return
	}

	// compare passwords
	if utils.ValidatePassword(user.Password, request.Password) != true {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Invalid password"})
		return
	}

	// gen token
	token, err := lc.UserUsecase.CreateToken(&user, lc.Env.TokenSecret, lc.Env.TokenTimeoutHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	response := domain.LoginResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}

func (lc *UserController) GetById(c *gin.Context) {
	user, err := lc.UserUsecase.GetById(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (lc *UserController) GetUsersByTeamId(c *gin.Context) {
	users, err := lc.UserUsecase.GetUsersByTeamId(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (lc *UserController) GetAll(c *gin.Context) {
	users, err := lc.UserUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
