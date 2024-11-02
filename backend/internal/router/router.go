package router

import (
	"backend/internal/config"
	handlers "backend/internal/handlers/common"
	userHandler "backend/internal/handlers/user"
	"backend/internal/middlewares"
	userRepositories "backend/internal/repositories/user"
	userServices "backend/internal/services/user"
	"backend/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, config *config.Config) *mux.Router {

	// Init repositories
	userRepo := userRepositories.NewUserRepo(db)
	userProfileRepo := userRepositories.NewUserProfileRepo(db)
	// ...

	// Init services
	userService := userServices.NewUserService(config, userRepo, userProfileRepo)
	// ...

	// Init handlers
	commonHandler := handlers.NewHealthHandler(config)
	userHandler := userHandler.NewUserHandler(userService)
	// ...

	r := mux.NewRouter()
	r = r.SkipClean(true) // todo: see what this does
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	// Unauthenticated routes
	r.HandleFunc("/health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/user/login", userHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/user/register-attendee", userHandler.RegisterAttendee).Methods(http.MethodPost)
	// ...

	protectedRouter := mux.NewRouter()
	protectedRouter = protectedRouter.SkipClean(true)
	protectedRouter.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)
	protectedRouter.Use(middlewares.Auth(utils.NewJWTUtil(config.JWT.Secret)))

	// Authenticated routes
	protectedRouter.HandleFunc("/secure-health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.PathPrefix("").Handler(protectedRouter)
	protectedRouter.HandleFunc("/user/profile", userHandler.CreateUserProfile).Methods(http.MethodPost)
	// ...

	return r
}
