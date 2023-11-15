package repository

import (
	"database/sql"
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

func (u *userRepository) Create(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetById(id int) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetByLogin(login string) (domain.User, error) {
	//table := u.table
	//rows, err := u.database.SelectAllFromX()
	// TODO
	panic("")
}

func (u *userRepository) GetUsersByTeamId(id int) ([]domain.User, error) {
	// todo
	panic("")
}

func (u *userRepository) GetAll() ([]domain.User, error) {
	// todo
	panic("")
}
