package router

import (
	"backend/internal/common/handlers"
	"backend/internal/config"
	"backend/internal/middlewares"
	"backend/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, config *config.Config) *mux.Router {

	// Init repositories
	// userRepo := repos.NewUserRepo(db)
	// ...

	// Init services
	// userService := services.NewUserService(userRepo, config)
	// ...

	// Init handlers
	commonHandler := handlers.NewCommonHandler(config)
	// userHandler := handlers.NewUserHandler(userService)
	// ...

	r := mux.NewRouter()
	r = r.SkipClean(true) // todo: see what this does
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	// Unauthenticated routes
	r.HandleFunc("/health", commonHandler.HealthCheck).Methods(http.MethodGet)
	// ...

	protectedRouter := mux.NewRouter()
	protectedRouter = protectedRouter.SkipClean(true)
	protectedRouter.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)
	protectedRouter.Use(middlewares.Auth(utils.NewJWTUtil(config.JWT.Secret)))

	// Authenticated routes
	protectedRouter.HandleFunc("/secure-health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.PathPrefix("").Handler(protectedRouter)
	// ...

	return r
}
