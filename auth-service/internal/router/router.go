package router

import (
	"auth-service/internal/config"
	"auth-service/internal/handlers"
	"auth-service/internal/repos"
	"auth-service/internal/services"
	"net/http"

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
	r.HandleFunc(config.App.BaseURL+"/register-attendee", userHandler.Create).Methods(http.MethodPost)
	// ...

	// Authenticated routes
	// ...

	return r
}
