package repository

import (
	"testGoLang/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(loginRequest model.LoginRequest) (model.User, error)
	Register(user model.User) (model.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) Login(loginRequest model.LoginRequest) (model.User, error) {
	var user model.User

	err := r.db.Find(&user, "email=?", loginRequest.Email).Error

	return user, err
}

func (r *authRepository) Register(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
