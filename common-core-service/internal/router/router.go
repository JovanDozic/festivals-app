package router

import (
	"common-core-service/internal/config"
	"common-core-service/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, config *config.Config) *mux.Router {

	// Init repositories
	// ...

	// Init services
	// ...

	// Init handlers
	commonHandler := handlers.NewCommonHandler(config)
	// ...

	r := mux.NewRouter()
	r = r.SkipClean(true) // todo: see what this does
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowedHandler)

	// Unauthenticated routes
	r.HandleFunc(config.App.BaseURL+"/health", commonHandler.HealthCheck).Methods(http.MethodGet)
	// ...

	// Authenticated routes
	// ...

	return r
}
