package services

import (
	"backend/internal/config"
	modelsError "backend/internal/models"
	modelsCommon "backend/internal/models/common"
	modelsUser "backend/internal/models/user"
	reposUser "backend/internal/repositories/user"
	servicesCommon "backend/internal/services/common"
	"backend/internal/utils"
	"log"
	"strings"
)

type UserService interface {
	Create(user *modelsUser.User) error
	Login(username string, password string) (string, error)
	CreateUserProfile(username string, userProfile *modelsUser.UserProfile) error
	CreateUserAddress(username string, address *modelsCommon.Address) error
}

type userService struct {
	config          *config.Config
	userRepo        reposUser.UserRepo
	profileRepo     reposUser.UserProfileRepo
	locationService servicesCommon.LocationService
}

func NewUserService(c *config.Config, r reposUser.UserRepo, p reposUser.UserProfileRepo, l servicesCommon.LocationService) UserService {
	return &userService{userRepo: r, config: c, profileRepo: p, locationService: l}
}

func (s *userService) Create(user *modelsUser.User) error {

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
			return modelsError.ErrDuplicateUser
		case strings.Contains(err.Error(), "foreign key constraint"):
			return modelsError.ErrRoleNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *userService) Login(username string, password string) (string, error) {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", modelsError.ErrNotFound
	}

	if !utils.VerifyPassword(user.Password, password) {
		return "", modelsError.ErrInvalidPassword
	}

	jwt := utils.NewJWTUtil(s.config.JWT.Secret)
	token, err := jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) CreateUserProfile(username string, userProfile *modelsUser.UserProfile) error {

	if err := userProfile.Validate(); err != nil {
		return err
	}

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return modelsError.ErrNotFound
	}

	userProfile.UserID = user.ID

	// * We don't need to check if user already has a profile because we are using a unique constraint on the user_id column

	if err := s.profileRepo.Create(userProfile); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return modelsError.ErrUserHasProfile
		case strings.Contains(err.Error(), "foreign key constraint"):
			return modelsError.ErrUserHasProfile
		default:
			return err
		}
	}

	return nil
}

func (s *userService) CreateUserAddress(username string, address *modelsCommon.Address) error {

	if err := address.Validate(); err != nil {
		log.Println("error validating address", err)
		return err
	}

	// todo: do we need to check if address already exists?
	// ! yes

	if err := s.locationService.CreateAddress(address); err != nil {
		log.Println("error creating address", err)
		return err
	}

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		log.Println("error getting user", err)
		return modelsError.ErrNotFound
	}

	err = s.profileRepo.UpdateAddressId(user.ID, address.ID)
	if err != nil {
		log.Println("error updating address id", err)
		return err
	}

	log.Println("user's address created successfully")
	return nil
}
