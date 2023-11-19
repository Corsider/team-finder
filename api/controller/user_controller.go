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

func (lc *UserController) RegUser(c *gin.Context) {
	var request domain.UserRegRequest
	err0 := c.BindJSON(&request)
	if err0 != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err0.Error()})
		return
	}

	count, err := lc.UserUsecase.CheckForExistence(request.Nickname, request.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	if count != 0 {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Error: "exist"})
		return
	}

	userId, err1 := lc.UserUsecase.InsertUser(request)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	userCreated, err2 := lc.UserUsecase.GetById(userId)
	if err2 != nil {
		lc.UserUsecase.DeleteUserById(userId) // check?
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	token, err3 := lc.UserUsecase.CreateToken(&userCreated, lc.Env.TokenSecret, lc.Env.TokenTimeoutHour)

	if err3 != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.UserRegResponse{UserId: strconv.Itoa(userId), Token: token})
}

func (lc *UserController) UpdateUser(c *gin.Context) {
	var request domain.UpdateRequest
	err0 := c.BindJSON(&request)
	if err0 != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err0.Error()})
		return
	}

	current, err := lc.UserUsecase.GetById(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	if request.Name == "" {
		request.Name = current.Name
	}
	if request.Nickname == "" {
		request.Nickname = current.Nickname
	}
	if request.Description == "" {
		request.Description = current.Description
	}

	user, err1 := lc.UserUsecase.UpdateUser(request, utils.First(strconv.Atoi(c.Param("id"))))
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err1.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
