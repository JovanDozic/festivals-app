package repositories

import (
	models "backend/internal/models/user"

	"gorm.io/gorm"
)

type UserProfileRepo interface {
	Create(userProfile *models.UserProfile) error
	GetByUserID(userID uint) (*models.UserProfile, error)
	GetFullByUsername(username string) (*models.UserProfile, error)
	UpdateAddressId(userID uint, addressID uint) error
	Update(userProfile *models.UserProfile) error
	GetFestivalEmployees(festivalId uint) ([]models.UserProfile, error)
	GetEmployeesNotOnFestival(festivalId uint) ([]models.UserProfile, error)
}

type userProfileRepo struct {
	db *gorm.DB
}

func NewUserProfileRepo(db *gorm.DB) UserProfileRepo {
	return &userProfileRepo{db}
}

func (r *userProfileRepo) Create(userProfile *models.UserProfile) error {
	return r.db.Create(userProfile).Error
}

func (r *userProfileRepo) GetByUserID(userID uint) (*models.UserProfile, error) {
	var userProfile models.UserProfile
	err := r.db.Where("user_id = ?", userID).First(&userProfile).Error
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

// Returns user profile object with joined all related data like address, city, country and image
func (r *userProfileRepo) GetFullByUsername(username string) (*models.UserProfile, error) {
	var userProfile models.UserProfile
	err := r.db.Preload("User").
		Preload("Address.City.Country").
		Preload("Image").
		Joins("LEFT JOIN users ON users.id = user_profiles.user_id").
		Where("username = ?", username).First(&userProfile).Error
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (r *userProfileRepo) UpdateAddressId(userID uint, addressID uint) error {
	return r.db.Model(&models.UserProfile{}).Where("user_id = ?", userID).Update("address_id", addressID).Error
}

func (r *userProfileRepo) Update(userProfile *models.UserProfile) error {
	return r.db.Save(userProfile).Error
}

func (r *userProfileRepo) GetFestivalEmployees(festivalId uint) ([]models.UserProfile, error) {
	var profiles []models.UserProfile

	err := r.db.Table("user_profiles").
		Preload("User").
		Preload("Address.City.Country").
		Joins("join users u on user_profiles.user_id = u.id").
		Joins("join employees e on u.id = e.user_id").
		Joins("join festival_employees fe on e.user_id = fe.user_id").
		Where("fe.festival_id = ?", festivalId).
		Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *userProfileRepo) GetEmployeesNotOnFestival(festivalId uint) ([]models.UserProfile, error) {
	var profiles []models.UserProfile

	err := r.db.Table("user_profiles").
		Preload("User").
		Preload("Address.City.Country").
		Joins("join users u on user_profiles.user_id = u.id").
		Joins("join employees e on u.id = e.user_id").
		Where("e.user_id not in (select user_id from festival_employees where festival_id = ?)", festivalId).
		Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	return profiles, nil
}
