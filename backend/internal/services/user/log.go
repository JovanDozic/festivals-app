package services

import (
	modelsUser "backend/internal/models/user"
	reposUser "backend/internal/repositories/user"
)

type LogService interface {
	GetAll() ([]modelsUser.Log, error)
	GetByUserRole(userRole modelsUser.UserRole) ([]modelsUser.Log, error)
}

type logService struct {
	logRepo reposUser.LogRepo
}

func NewLogService(lr reposUser.LogRepo) LogService {
	return &logService{
		logRepo: lr,
	}
}

func (ls *logService) GetAll() ([]modelsUser.Log, error) {
	return ls.logRepo.GetAll()
}

func (ls *logService) GetByUserRole(userRole modelsUser.UserRole) ([]modelsUser.Log, error) {
	return ls.logRepo.GetAllByRole(userRole)
}
