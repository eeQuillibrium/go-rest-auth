package service

import (
	auth "github.com/eeQuillibrium/go-rest-auth"
	"github.com/eeQuillibrium/go-rest-auth/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth.User) (int, error)
	CheckUser(user auth.User) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos.Authorization)}
}
