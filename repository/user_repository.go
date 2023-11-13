package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"team-finder/domain"
	"team-finder/postgres"
)

type userRepository struct {
	database postgres.Database //*sql.DB
	table    string
}

func NewUserRepository(db *sql.DB, table string) domain.UserRepository {
	return &userRepository{
		database: &postgres.PostgresDB{DB: db},
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
	//table := u.table
	//rows, err := u.database.SelectAllFromX()
	// TODO
	panic("")
}
