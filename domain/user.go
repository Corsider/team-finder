package domain

import "github.com/gin-gonic/gin"

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
	Create(c *gin.Context, user *User) error
	GetById(c *gin.Context) (User, error)
	GetByLogin(c *gin.Context, login string) (User, error)
}
