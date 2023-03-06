package service

import (
	"github.com/cucumberjaye/softline/internal/models"
)

type AuthRepository interface {
	GetUser(user models.LoginUser) (models.User, error)
	CreateUser(user models.RegisterUser) (int, error)
}

type Service struct {
	authRepository AuthRepository
}

func New(authRepo AuthRepository) *Service {
	return &Service{
		authRepository: authRepo,
	}
}
