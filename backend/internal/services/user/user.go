package services

import (
	"backend/internal/config"
	dtoCommon "backend/internal/dto/common"
	dto "backend/internal/dto/user"
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
	CreateAttendee(user *modelsUser.User) error
	Login(username string, password string) (string, error)
	GetUserProfile(username string) (*dto.GetUserProfileResponse, error)
	CreateUserProfile(username string, userProfile *modelsUser.UserProfile) error
	CreateUserAddress(username string, address *modelsCommon.Address) error
	ChangePassword(username, oldPassword, newPassword string) error
	UpdateUserProfile(username string, updatedProfile *modelsUser.UserProfile) error
	UpdateUserEmail(username string, email string) error
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

func (s *userService) CreateAttendee(user *modelsUser.User) error {

	if err := user.Validate(); err != nil {
		return err
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash

	if err := s.userRepo.CreateAttendee(user); err != nil {
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

func (s *userService) GetUserProfile(username string) (*dto.GetUserProfileResponse, error) {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, modelsError.ErrNotFound
	}

	userProfile, err := s.profileRepo.GetFullByUsername(username)
	if err != nil {
		return nil, modelsError.ErrNotFound
	}

	response := dto.GetUserProfileResponse{
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role,
		FirstName:   userProfile.FirstName,
		LastName:    userProfile.LastName,
		DateOfBirth: userProfile.DateOfBirth.Format("2006-01-02"),
		PhoneNumber: userProfile.PhoneNumber,
		Address:     nil,
	}

	if userProfile.Address != nil {
		response.Address = &dtoCommon.GetAddressResponse{
			Street:         userProfile.Address.Street,
			Number:         userProfile.Address.Number,
			ApartmentSuite: userProfile.Address.ApartmentSuite,
			City:           userProfile.Address.City.Name,
			PostalCode:     userProfile.Address.City.PostalCode,
			Country:        userProfile.Address.City.Country.NiceName,
			CountryISO3:    userProfile.Address.City.Country.ISO3,
			CountryISO2:    userProfile.Address.City.Country.ISO,
		}
	}

	return &response, nil
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

func (s *userService) ChangePassword(username, oldPassword, newPassword string) error {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return modelsError.ErrNotFound
	}

	if !utils.VerifyPassword(user.Password, oldPassword) {
		return modelsError.ErrInvalidPassword
	}

	passwordHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	err = s.userRepo.UpdatePassword(username, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) CreateUserAddress(username string, address *modelsCommon.Address) error {

	if err := address.Validate(); err != nil {
		log.Println("error validating address", err)
		return err
	}

	userProfile, err := s.profileRepo.GetFullByUsername(username)
	if err != nil {
		log.Println("error getting user profile", err)
		return modelsError.ErrNotFound
	}

	if userProfile.AddressID != nil {
		log.Println("user already has an address")
		return modelsError.ErrUserHasAddress
	}

	if err := s.locationService.CreateAddress(address); err != nil {
		log.Println("error creating address", err)
		return err
	}

	err = s.profileRepo.UpdateAddressId(userProfile.UserID, address.ID)
	if err != nil {
		log.Println("error updating address id", err)
		return err
	}

	log.Println("user's address created successfully")
	return nil
}

func (s *userService) UpdateUserEmail(username string, email string) error {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return modelsError.ErrNotFound
	}

	user.Email = email

	if err := s.userRepo.Update(user); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return modelsError.ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (s *userService) UpdateUserProfile(username string, updatedProfile *modelsUser.UserProfile) error {

	if err := updatedProfile.Validate(); err != nil {
		return err
	}

	profile, err := s.profileRepo.GetFullByUsername(username)
	if err != nil {
		return modelsError.ErrNotFound
	}

	profile.FirstName = updatedProfile.FirstName
	profile.LastName = updatedProfile.LastName
	profile.DateOfBirth = updatedProfile.DateOfBirth
	profile.PhoneNumber = updatedProfile.PhoneNumber

	if err := s.profileRepo.Update(profile); err != nil {
		return err
	}

	return nil
}
