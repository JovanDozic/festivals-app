package services

import (
	"backend/internal/config"
	errorModels "backend/internal/models"
	models "backend/internal/models/user"
	repositories "backend/internal/repositories/user"
	"backend/internal/utils"
	"strings"
)

type UserService interface {
	Create(user *models.User) error
	Login(username string, password string) (string, error)
	CreateUserProfile(username string, userProfile *models.UserProfile) error
}

type userService struct {
	userRepo    repositories.UserRepo
	profileRepo repositories.UserProfileRepo
	config      *config.Config
}

func NewUserService(c *config.Config, r repositories.UserRepo, p repositories.UserProfileRepo) UserService {
	return &userService{userRepo: r, config: c, profileRepo: p}
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

	if err := s.userRepo.Create(user); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return errorModels.ErrDuplicateUser
		case strings.Contains(err.Error(), "foreign key constraint"):
			return errorModels.ErrRoleNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *userService) Login(username string, password string) (string, error) {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", errorModels.ErrNotFound
	}

	if !utils.VerifyPassword(user.Password, password) {
		return "", errorModels.ErrInvalidPassword
	}

	jwt := utils.NewJWTUtil(s.config.JWT.Secret)
	token, err := jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) CreateUserProfile(username string, userProfile *models.UserProfile) error {

	if err := userProfile.Validate(); err != nil {
		return err
	}

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return errorModels.ErrNotFound
	}

	userProfile.UserID = user.UserID

	// ? We don't need to check if user already has a profile because we are using a unique constraint on the user_id column

	if err := s.profileRepo.Create(userProfile); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return errorModels.ErrUserHasProfile
		case strings.Contains(err.Error(), "foreign key constraint"):
			return errorModels.ErrUserHasProfile
		default:
			return err
		}
	}

	return nil
}
