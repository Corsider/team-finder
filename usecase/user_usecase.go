package usecase

import (
	"team-finder/domain"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *userUsecase) GetUserByLogin(login string) (domain.User, error) {
	u, err := lu.userRepository.GetByLogin(login)
	return u, err
}

func (lu *userUsecase) CreateToken(user *domain.User, secret string, exp int) (token string, err error) {
	//TODO implement me
	panic("implement me")
}

func (lu *userUsecase) GetById(id int) (domain.User, error) {
	return lu.userRepository.GetById(id)
}

func (lu *userUsecase) GetUsersByTeamId(id int) ([]domain.User, error) {
	return lu.userRepository.GetUsersByTeamId(id)
}

func (lu *userUsecase) GetAll() ([]domain.User, error) {
	return lu.userRepository.GetAll()
}
