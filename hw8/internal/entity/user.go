package entity

import "github.com/golang-jwt/jwt"

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type Token struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	*jwt.StandardClaims
}
