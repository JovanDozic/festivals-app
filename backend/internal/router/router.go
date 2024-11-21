package router

import (
	"backend/internal/config"
	handlersCommon "backend/internal/handlers/common"
	handlers "backend/internal/handlers/festival"
	handlersUser "backend/internal/handlers/user"
	"backend/internal/middlewares"
	reposCommon "backend/internal/repositories/common"
	reposFestival "backend/internal/repositories/festival"
	reposUser "backend/internal/repositories/user"
	servicesCommon "backend/internal/services/common"
	services "backend/internal/services/festival"
	servicesUser "backend/internal/services/user"
	"backend/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, config *config.Config) *mux.Router {

	// Init repositories
	userRepo := reposUser.NewUserRepo(db)
	userProfileRepo := reposUser.NewUserProfileRepo(db)
	addressRepo := reposCommon.NewAddressRepo(db)
	cityRepo := reposCommon.NewCityRepo(db)
	countryRepo := reposCommon.NewCountryRepo(db)
	festivalRepo := reposFestival.NewFestivalRepo(db)
	imageRepo := reposCommon.NewImageRepo(db)
	// ...

	// Init services
	locationService := servicesCommon.NewLocationService(addressRepo, cityRepo, countryRepo)
	userService := servicesUser.NewUserService(config, userRepo, userProfileRepo, locationService)
	festivalService := services.NewFestivalService(config, festivalRepo, userRepo, locationService, imageRepo)
	// ...

	// Init handlers
	commonHandler := handlersCommon.NewHealthHandler(config)
	userHandler := handlersUser.NewUserHandler(userService)
	festivalHandler := handlers.NewFestivalHandler(festivalService, locationService)
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
	protectedRouter.HandleFunc("/user/profile/address", userHandler.CreateUserAddress).Methods(http.MethodPost)
	protectedRouter.HandleFunc("/user/profile", userHandler.GetUserProfile).Methods(http.MethodGet)
	protectedRouter.HandleFunc("/user/change-password", userHandler.ChangePassword).Methods(http.MethodPut)
	protectedRouter.HandleFunc("/user/profile", userHandler.UpdateUserProfile).Methods(http.MethodPut)
	protectedRouter.HandleFunc("/user/email", userHandler.UpdateUserEmail).Methods(http.MethodPut)
	// ...
	// todo: should this be like get all future ones, or ones in the attendee's city?
	protectedRouter.HandleFunc("/festival", festivalHandler.GetAll).Methods(http.MethodGet)
	// ... ORGANIZER ONLY
	protectedRouter.HandleFunc("/festival", festivalHandler.Create).Methods(http.MethodPost)
	protectedRouter.HandleFunc("/organizer/festival", festivalHandler.GetByOrganizer).Methods(http.MethodGet)
	protectedRouter.HandleFunc("/festival/{festivalId}", festivalHandler.GetById).Methods(http.MethodGet)
	protectedRouter.HandleFunc("/festival/{festivalId}", festivalHandler.Update).Methods(http.MethodPut)
	protectedRouter.HandleFunc("/festival/{festivalId}", festivalHandler.Delete).Methods(http.MethodDelete)

	protectedRouter.HandleFunc("/festival/{festivalId}/publish", festivalHandler.PublishFestival).Methods(http.MethodPut)
	protectedRouter.HandleFunc("/festival/{festivalId}/cancel", festivalHandler.CancelFestival).Methods(http.MethodPut)
	protectedRouter.HandleFunc("/festival/{festivalId}/complete", festivalHandler.CompleteFestival).Methods(http.MethodPut)
	protectedRouter.HandleFunc("/festival/{festivalId}/store/open", festivalHandler.OpenStore).Methods(http.MethodPut)
	protectedRouter.HandleFunc("/festival/{festivalId}/store/close", festivalHandler.CloseStore).Methods(http.MethodPut)

	protectedRouter.HandleFunc("/festival/{festivalId}/image", festivalHandler.GetImages).Methods(http.MethodGet)
	// ? Should we return list of images in get festival by id?
	protectedRouter.HandleFunc("/festival/{festivalId}/image", festivalHandler.AddImage).Methods(http.MethodPost)

	// ...

	return r
}
