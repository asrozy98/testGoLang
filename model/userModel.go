package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"password"`
	// Token    string `gorm:"-"`
}

type RegisterRequest struct {
	// gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	// Token    string `gorm:"-"`
}

type LoginRequest struct {
	// gorm.Model
	// Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	// Token    string `gorm:"-"`
}