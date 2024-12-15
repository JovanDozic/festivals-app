package services

import (
	models "backend/internal/models/user"
	repositories "backend/internal/repositories/common"
)

type LogService interface {
	GetAll() ([]models.Log, error)
	GetByUserRole(userRole models.UserRole) ([]models.Log, error)
}

type logService struct {
	logRepo repositories.LogRepo
}

func NewLogService(lr repositories.LogRepo) LogService {
	return &logService{
		logRepo: lr,
	}
}

func (ls *logService) GetAll() ([]models.Log, error) {
	return ls.logRepo.GetAll()
}

func (ls *logService) GetByUserRole(userRole models.UserRole) ([]models.Log, error) {
	return ls.logRepo.GetAllByRole(userRole)
}
