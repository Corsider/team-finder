package repository

import (
	"database/sql"
	"strconv"
	"team-finder/domain"
	"team-finder/internal/utils"
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

func (u *userRepository) CheckForExistence(nickname, login string) (int, error) {
	count, err := u.database.SelectCountFromXWhereYeqZorNeqM("users", "nickname", "login", nickname, login)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (u *userRepository) InsertUser(request domain.UserRegRequest) (int, error) {
	userId, err := u.database.InsertParametrizedIntoXYValuesZReturningN(u.table, "name, nickname, rate, description, login, password",
		"$1, $2, 5, $3, $4, $5", "user_id", request.Name, request.Nickname, request.Description, request.Login,
		utils.HashPassword(request.Password))
	if err != nil {
		return 0, err
	}
	return utils.First(strconv.Atoi(strconv.FormatInt(userId.(int64), 10))), nil
}

func (u *userRepository) DeleteUserById(userId int) error {
	err := u.database.DeleteFromXWhereYeqZ(u.table, "user_id", strconv.Itoa(userId))
	return err
}

func (u *userRepository) UpdateUser(request domain.UpdateRequest, userId int) (domain.User, error) {
	entity, err := u.database.UpdateNXSetYZWhereNeqM("users", "user_id", strconv.Itoa(userId), []string{"name", "nickname", "description"}, domain.User{}, request.Name, request.Nickname, request.Description)
	if err != nil {
		return domain.User{}, err
	}
	usr, _ := entity.(domain.User)
	return usr, nil
}
