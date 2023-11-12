package usecase

import (
	"github.com/gin-gonic/gin"
	"team-finder/domain"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByLogin(c *gin.Context, login string) (domain.User, error) {
	u, err := lu.userRepository.GetByLogin(c, login)
	return u, err
}

func (lu *loginUsecase) CreateToken(user *domain.User, secret string, exp int) (token string, err error) {
	//TODO implement me
	panic("implement me")
}
