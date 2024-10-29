package router

import (
	"net/http"
	"user-service/internal/config"
	"user-service/internal/handlers"
	"user-service/internal/middlewares"
	"user-service/internal/repos"
	"user-service/internal/services"
	"user-service/internal/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, config *config.Config) *mux.Router {

	// Init repositories
	userRepo := repos.NewUserRepo(db)
	// ...

	// Init services
	userService := services.NewUserService(userRepo, config)
	// ...

	// Init handlers
	commonHandler := handlers.NewCommonHandler(config)
	userHandler := handlers.NewUserHandler(userService)
	// ...

	r := mux.NewRouter()
	r = r.SkipClean(true) // todo: see what this does
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowedHandler)

	// Unauthenticated routes
	r.HandleFunc(config.App.BaseURL+"/health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc(config.App.BaseURL+"/register-attendee", userHandler.RegisterAttendee).Methods(http.MethodPost)
	r.HandleFunc(config.App.BaseURL+"/login", userHandler.Login).Methods(http.MethodPost)
	// todo: temp:
	r.HandleFunc(config.App.BaseURL+"/try-address", userHandler.TryAddress).Methods(http.MethodPost)
	// ...

	protectedRouter := mux.NewRouter()
	protectedRouter = protectedRouter.SkipClean(true)
	protectedRouter.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowedHandler)
	protectedRouter.Use(middlewares.Auth(utils.NewJWTUtil(config.JWT.Secret)))

	// Authenticated routes
	protectedRouter.HandleFunc(config.App.BaseURL+"/secure-health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.PathPrefix(config.App.BaseURL).Handler(protectedRouter)
	// ...

	return r
}
