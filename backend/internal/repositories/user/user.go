package repositories

import (
	models "backend/internal/models/user"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *models.User) error
	CreateAttendee(user *models.User) error
	GetByUsername(username string) (*models.User, error)
	GetById(id uint) (*models.User, error)
	UpdatePassword(username, password string) error
	Update(user *models.User) error
	CreateEmployee(user *models.User) error
	CreateOrganizer(user *models.User) error
	CreateAdmin(user *models.User) error
	GetAttendeeCount() (int, error)
	GetIdByUsername(username string) (uint, error)
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

func (r *userRepo) GetById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
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

	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	attendee := &models.Attendee{
		UserID: user.ID,
		User:   *user,
	}

	err = r.db.Create(attendee).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) CreateEmployee(user *models.User) error {

	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	employee := &models.Employee{
		UserID: user.ID,
		User:   *user,
	}

	err = r.db.Create(employee).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) CreateOrganizer(user *models.User) error {

	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	employee := &models.Organizer{
		UserID: user.ID,
		User:   *user,
	}

	err = r.db.Create(employee).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) CreateAdmin(user *models.User) error {

	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	employee := &models.Administrator{
		UserID: user.ID,
		User:   *user,
	}

	err = r.db.Create(employee).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetAttendeeCount() (int, error) {
	var count int64
	err := r.db.Model(&models.Attendee{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *userRepo) GetIdByUsername(username string) (uint, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
