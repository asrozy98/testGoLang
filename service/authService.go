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
	Register(registerRequest model.RegisterRequest) (model.User, error)
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
	// var w http.ResponseWriter
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

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &model.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte("my_secret_key")
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "Internal Server Error", err
	}

	return tokenString, err
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
