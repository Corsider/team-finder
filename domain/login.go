package domain

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Login    string `form:"login" binding:"required,login"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginUsecase interface {
	GetUserByLogin(c *gin.Context, login string) (User, error)
	CreateToken(user *User, secret string, exp int) (token string, err error)
}
