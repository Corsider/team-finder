package usecase

import (
	"team-finder/domain"
	"team-finder/internal/utils"
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
	return utils.CreateToken(user, secret, exp)
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

func (lu *userUsecase) CheckForExistence(nickname, login string) (int, error) {
	return lu.userRepository.CheckForExistence(nickname, login)
}

func (lu *userUsecase) InsertUser(request domain.UserRegRequest) (int, error) {
	return lu.userRepository.InsertUser(request)
}

func (lu *userUsecase) DeleteUserById(userId int) error {
	return lu.userRepository.DeleteUserById(userId)
}
