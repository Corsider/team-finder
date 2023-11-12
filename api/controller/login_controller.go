package controller

import (
	"github.com/gin-gonic/gin"
	"team-finder/boot"
	"team-finder/domain"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *boot.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	// TODO
}
