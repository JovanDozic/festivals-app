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
	Login(username string, password string) (string, error)
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
			return models.ErrDuplicateUser
		case strings.Contains(err.Error(), "foreign key constraint"):
			return models.ErrRoleNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *userService) Login(username string, password string) (string, error) {

	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", models.ErrNotFound
	}

	if !utils.VerifyPassword(user.Password, password) {
		return "", models.ErrInvalidPassword
	}

	jwt := utils.NewJWTUtil(s.config.JWT.Secret)
	token, err := jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
