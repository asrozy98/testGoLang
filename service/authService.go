package service

import (
	"net/http"
	"testGoLang/model"
	"testGoLang/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(loginRequest model.LoginRequest) (string, error)
	Register(registerRequest model.RegisterRequest) (any, error)
}

type authService struct {
	authRepository repository.AuthRepository
}
type authCookie struct {
	w http.ResponseWriter
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
		return "Password failure", err
	}

	expirationTime := time.Now().Add(10 * time.Minute)

	claims := &model.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(model.JwtKey)
	if err != nil {
		return "Internal Server Error", err
	}

	return tokenString, err
}

func (s *authService) Register(registerRequest model.RegisterRequest) (any, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	user := model.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: password,
	}
	result, err := s.authRepository.Register(user)
	if err != nil {
		return "Bad Request", err
	}

	return result, err
}
