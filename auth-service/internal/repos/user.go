package repos

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *models.User) error
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
