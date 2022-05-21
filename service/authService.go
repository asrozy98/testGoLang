package service

import (
	"testGoLang/model"
	"testGoLang/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(loginRequest model.LoginRequest) (string, error)
	Register(registerRequest model.RegisterRequest) (model.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository}
}

func (s *authService) Login(loginRequest model.LoginRequest) (string, error) {
	user, err := s.authRepository.Login(loginRequest)

	if err != nil {
		return "Bad Request", err
	}

	if user.ID == 0 {
		return "User Not Found", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return "Password failur", err
	}

	return "Login Success", err
}

func (s *authService) Register(registerRequest model.RegisterRequest) (model.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	user := model.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: password,
	}
	newUser, err := s.authRepository.Register(user)
	return newUser, err
}
