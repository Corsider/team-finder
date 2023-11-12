package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"team-finder/domain"
)

type userRepository struct {
	database *sql.DB
	table    string
}

func NewUserRepository(db *sql.DB, table string) domain.UserRepository {
	return &userRepository{
		database: db,
		table:    table,
	}
}

func (u *userRepository) Create(c *gin.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetById(c *gin.Context) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetByLogin(c *gin.Context, login string) (domain.User, error) {
	//todo
	return domain.User{}, nil
}
