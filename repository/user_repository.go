package repository

import (
	"database/sql"
	"strconv"
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
	row := u.database.Select1FromXWhereYeqZ(u.table, "user_id", strconv.Itoa(id))
	var usr domain.User
	err := row.Scan(&usr.UserId, &usr.Name, &usr.Nickname, &usr.Rate, &usr.Description, &usr.Login, &usr.Password)
	if err != nil {
		return domain.User{}, err
	}
	return usr, nil
}

func (u *userRepository) GetByLogin(login string) (domain.User, error) {
	//rows, err := u.database.SelectAllFromX()
	panic("inop")
}

func (u *userRepository) GetUsersByTeamId(id int) ([]domain.User, error) {
	rows, err := u.database.SelectAllFromXWhereYeqZ("user_team", "team_id", strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	users := []domain.User{}
	for rows.Next() {
		var usr domain.User
		rows.Scan(&usr.UserId, &usr.Name, &usr.Nickname, &usr.Rate, &usr.Description, &usr.Login, &usr.Password)
		usr.Password = ""
		users = append(users, usr)
	}
	return users, nil
}

func (u *userRepository) GetAll() ([]domain.User, error) {
	rows, err := u.database.SelectAllFromX(u.table)
	if err != nil {
		return nil, err
	}
	users := []domain.User{}
	for rows.Next() {
		var usr domain.User
		rows.Scan(&usr.UserId, &usr.Name, &usr.Nickname, &usr.Rate, &usr.Description, &usr.Login, &usr.Password)
		usr.Password = ""
		users = append(users, usr)
	}
	return users, nil
}
