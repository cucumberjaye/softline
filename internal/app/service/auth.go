package service

import (
	"github.com/cucumberjaye/softline/configs"
	"github.com/cucumberjaye/softline/internal/models"
	"github.com/cucumberjaye/softline/pkg/hasher"
	"github.com/cucumberjaye/softline/pkg/jwt_manager"
)

func (s *Service) CreateUser(user models.RegisterUser) (int, error) {
	user.Password = hasher.GeneratePasswordHash(user.Password)
	return s.authRepository.CreateUser(user)
}

func (s *Service) GenerateToken(loginUser models.LoginUser) (string, error) {
	loginUser.Password = hasher.GeneratePasswordHash(loginUser.Password)
	user, err := s.authRepository.GetUser(loginUser)
	if err != nil {
		return "", err
	}

	return jwt_manager.GenerateToken(user.Id, []byte(configs.SigningKey))
}
