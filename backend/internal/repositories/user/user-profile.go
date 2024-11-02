package repositories

import (
	models "backend/internal/models/user"

	"gorm.io/gorm"
)

type UserProfileRepo interface {
	Create(userProfile *models.UserProfile) error
	GetByUserID(userID uint) (*models.UserProfile, error)
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
