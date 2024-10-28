package services

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"auth-service/internal/repos"
	"auth-service/internal/utils"
	"strings"
)

type UserService interface {
	Create(user *models.User) error
	// todo: other methods
}

type userService struct {
	repo   repos.UserRepo
	config *config.Config
}

func NewUserService(r repos.UserRepo, c *config.Config) UserService {
	return &userService{repo: r, config: c}
}

func (s *userService) Create(user *models.User) error {

	if err := user.Validate(); err != nil {
		return err
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash

	if err := s.repo.Create(user); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return models.ErrDuplicateUsername
		case strings.Contains(err.Error(), "foreign key constraint"):
			return models.ErrRoleNotFound
		default:
			return err
		}
	}

	return nil
}
