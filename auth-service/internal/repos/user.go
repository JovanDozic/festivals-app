package repos

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *models.User) error
	GetByUsername(username string) (*models.User, error)
	// todo: other methods
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
