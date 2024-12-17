package user

import (
	modelsUser "backend/internal/models/user"

	"gorm.io/gorm"
)

type LogRepo interface {
	CreateLog(log *modelsUser.Log) error
	GetAll() ([]modelsUser.Log, error)
	GetAllByRole(userRole modelsUser.UserRole) ([]modelsUser.Log, error)
	GetAllByRoles(userRoles []modelsUser.UserRole) ([]modelsUser.Log, error)
}

type logRepo struct {
	db *gorm.DB
}

func NewLogRepo(db *gorm.DB) LogRepo {
	return &logRepo{
		db: db,
	}
}

func (r *logRepo) CreateLog(log *modelsUser.Log) error {
	return r.db.Create(log).Error
}

func (r *logRepo) GetAll() ([]modelsUser.Log, error) {
	var logs []modelsUser.Log
	if err := r.db.
		Preload("User").
		Order("created_at DESC").
		Find(&logs).
		Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *logRepo) GetAllByRole(userRole modelsUser.UserRole) ([]modelsUser.Log, error) {
	var logs []modelsUser.Log

	if err := r.db.
		Preload("User").
		Joins("JOIN users ON logs.user_id = users.id").
		Where("users.role = ?", userRole).
		Order("created_at DESC").
		Find(&logs).
		Error; err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *logRepo) GetAllByRoles(userRoles []modelsUser.UserRole) ([]modelsUser.Log, error) {
	var logs []modelsUser.Log

	if err := r.db.
		Preload("User").
		Joins("JOIN users ON logs.user_id = users.id").
		Where("users.role IN (?)", userRoles).
		Order("created_at DESC").
		Find(&logs).
		Error; err != nil {
		return nil, err
	}

	return logs, nil
}
