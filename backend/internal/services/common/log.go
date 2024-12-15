package services

import (
	models "backend/internal/models/user"
	repositoriesCommon "backend/internal/repositories/common"
	repositories "backend/internal/repositories/user"
	"backend/internal/utils"
	"context"
	"log"
)

type LogService interface {
	Info(msg string, ctx context.Context)
	Error(msg string, ctx context.Context)
}

type logService struct {
	logRepo  repositoriesCommon.LogRepo
	userRepo repositories.UserRepo
}

func NewLogService(lr repositoriesCommon.LogRepo, ur repositories.UserRepo) LogService {
	return &logService{
		logRepo:  lr,
		userRepo: ur,
	}
}

func (s *logService) Info(msg string, ctx context.Context) {

	logModel := models.Log{
		Type:        "INFO",
		Description: msg,
	}

	username := utils.GetUsername(ctx)
	if username != "" {
		userId, err := s.userRepo.GetIdByUsername(username)
		if err != nil {
			log.Println("error saving log", err)
		} else {
			logModel.UserID = &userId
		}
		log.Println(msg + " (by: " + username + ")")
	} else {
		log.Println(msg)
	}

	err := s.logRepo.CreateLog(&logModel)
	if err != nil {
		log.Println("error saving log", err)
	}
}

func (s *logService) Error(msg string, ctx context.Context) {

	log.Print(msg)

	logModel := models.Log{
		Type:        "ERROR",
		Description: msg,
	}

	username := utils.GetUsername(ctx)
	if username != "" {
		userId, err := s.userRepo.GetIdByUsername(username)
		if err != nil {
			log.Println("error saving log", err)
		} else {
			logModel.UserID = &userId
		}
		log.Println(msg + " (by: " + username + ")")
	} else {
		log.Println(msg)
	}

	err := s.logRepo.CreateLog(&logModel)
	if err != nil {
		log.Println("error saving log", err)
	}
}
