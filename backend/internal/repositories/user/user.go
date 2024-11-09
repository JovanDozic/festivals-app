package repositories

import (
	models "backend/internal/models/user"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *models.User) error
	CreateAttendee(user *models.User) error
	GetByUsername(username string) (*models.User, error)
	UpdatePassword(username, password string) error
	Update(user *models.User) error
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

func (r *userRepo) UpdatePassword(username, password string) error {
	return r.db.Model(&models.User{}).Where("username = ?", username).Update("password", password).Error
}

func (r *userRepo) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) CreateAttendee(user *models.User) error {

	attendee := &models.Attendee{
		UserID: user.ID,
		User:   *user,
	}

	err := r.db.Create(attendee).Error
	if err != nil {
		return err
	}

	return nil
}
