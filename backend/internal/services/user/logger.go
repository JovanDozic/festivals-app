package services

import (
	modelsUser "backend/internal/models/user"
	reposUser "backend/internal/repositories/user"
	"backend/internal/utils"
	"context"
	"log"
)

type Logger interface {
	Info(msg string, ctx context.Context)
	Error(msg string, ctx context.Context)
}

type logger struct {
	logRepo  reposUser.LogRepo
	userRepo reposUser.UserRepo
}

func NewLogger(lr reposUser.LogRepo, ur reposUser.UserRepo) Logger {
	return &logger{
		logRepo:  lr,
		userRepo: ur,
	}
}

func (s *logger) Info(msg string, ctx context.Context) {

	logModel := modelsUser.Log{
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

func (s *logger) Error(msg string, ctx context.Context) {

	log.Print(msg)

	logModel := modelsUser.Log{
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
