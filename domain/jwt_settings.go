package domain

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	Login string `json:"login"`
	ID    string `json:"user_id"`
	jwt.StandardClaims
}
