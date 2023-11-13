package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"team-finder/boot"
	"team-finder/domain"
	"team-finder/internal/utils"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *boot.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	user, err := lc.LoginUsecase.GetUserByLogin(c, request.Login)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Not found"})
		return
	}

	// compare passwords
	if utils.ValidatePassword(user.Password, request.Password) != true {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid password"})
		return
	}

	// gen token
	token, err := lc.LoginUsecase.CreateToken(&user, lc.Env.TokenSecret, lc.Env.TokenTimeoutHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	response := domain.LoginResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}
