package repositories

import (
	models "backend/internal/models/user"

	"gorm.io/gorm"
)

type UserProfileRepo interface {
	Create(userProfile *models.UserProfile) error
	GetByUserID(userID uint) (*models.UserProfile, error)
	GetByUsername(username string) (*models.UserProfile, error)
	UpdateAddressId(userID uint, addressID uint) error
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

func (r *userProfileRepo) GetByUsername(username string) (*models.UserProfile, error) {
	var userProfile models.UserProfile
	err := r.db.Joins("JOIN users ON user_profiles.user_id = users.user_id").Where("users.username = ?", username).First(&userProfile).Error
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (r *userProfileRepo) UpdateAddressId(userID uint, addressID uint) error {
	return r.db.Model(&models.UserProfile{}).Where("user_id = ?", userID).Update("address_id", addressID).Error
}
