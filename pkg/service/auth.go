package service

import (
	"crypto/sha256"
	"fmt"
	"time"

	auth "github.com/eeQuillibrium/go-rest-auth"
	repository "github.com/eeQuillibrium/go-rest-auth/pkg/repository"
	"github.com/golang-jwt/jwt"
)

const (
	salt       string = "lrghl;dfjglj234"
	jwtSignKey string = "asdsajlfasd213;s"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user auth.User) (int, error) {
	user.Password = generatePassword(user.Password)
	return a.repo.CreateUser(user)
}
func (a *AuthService) CheckUser(user auth.User) (string, error) {
	user.Password = generatePassword(user.Password)

	id, err := a.repo.CheckUser(user)
	if err != nil {
		fmt.Println("user error id:", id)
		return "", err
	}

	token, err := createToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
func createToken(user auth.User) (string, error) {
	payload := tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Subject:   user.Login,
		},
		user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &payload)

	t, err := token.SignedString([]byte(jwtSignKey))
	if err != nil {

		return "", err
	}

	return t, nil
}
func generatePassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum([]byte(salt)))
}
