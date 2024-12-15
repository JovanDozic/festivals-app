package repositories

import (
	modelsUser "backend/internal/models/user"

	"gorm.io/gorm"
)

type LogRepo interface {
	CreateLog(log *modelsUser.Log) error
	GetLogs() ([]modelsUser.Log, error)
	GetLogsByRole(userRole modelsUser.UserRoles) ([]modelsUser.Log, error)
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

func (r *logRepo) GetLogs() ([]modelsUser.Log, error) {
	var logs []modelsUser.Log
	if err := r.db.
		Preload("User").
		Find(&logs).
		Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *logRepo) GetLogsByRole(userRole modelsUser.UserRoles) ([]modelsUser.Log, error) {
	var logs []modelsUser.Log

	if err := r.db.
		Preload("User").
		Joins("JOIN users ON logs.user_id = users.id").
		Where("users.role = ?", userRole).
		Find(&logs).
		Error; err != nil {
		return nil, err
	}

	return logs, nil
}
