package domain

const (
	TableUser = "users"
)

type User struct {
	UserId      int     `db:"user_id" json:"user_id"`
	Name        string  `db:"name" json:"name"`
	Nickname    string  `db:"nickname" json:"nickname"`
	Rate        float64 `db:"rate" json:"rate"`
	Description string  `db:"description" json:"description"`
	Login       string  `db:"login" json:"login"`
	Password    string  `db:"password" json:"password"`
}

type UserRepository interface {
	Create(user *User) error
	GetById(id int) (User, error)
	GetByLogin(login string) (User, error)
	GetUsersByTeamId(id int) ([]User, error)
	GetAll() ([]User, error)
	CheckForExistence(nickname, login string) (int, error)
	InsertUser(request UserRegRequest) (int, error)
	DeleteUserById(userId int) error
	UpdateUser(request UpdateRequest, userId int) (User, error)
}

type LoginRequest struct {
	Login    string `form:"login" binding:"required,login"`
	Password string `form:"password" binding:"required"`
}

type UpdateRequest struct {
	Name        string `form:"name"`
	Nickname    string `form:"nickname"`
	Description string `form:"description"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserRegRequest struct {
	Name        string `form:"name" binding:"required"`
	Nickname    string `form:"nickname" binding:"required"`
	Description string `form:"description" binding:"required"`
	Login       string `form:"login" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

type UserRegResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	User User `json:"user"`
}

type UsersResponse struct {
	Users []User `json:"users"`
}

type UserUsecase interface {
	GetUserByLogin(login string) (User, error)
	CreateToken(user *User, secret string, exp int) (token string, err error)
	GetById(id int) (User, error)
	GetUsersByTeamId(id int) ([]User, error)
	GetAll() ([]User, error)
	CheckForExistence(nickname, login string) (int, error)
	InsertUser(request UserRegRequest) (int, error)
	DeleteUserById(userId int) error
	UpdateUser(request UpdateRequest, userId int) (User, error)
}
