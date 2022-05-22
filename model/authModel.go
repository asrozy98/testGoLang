package model

import "github.com/dgrijalva/jwt-go"

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var JwtKey = []byte("my_secret_key")
