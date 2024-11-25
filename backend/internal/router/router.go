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

	protR := mux.NewRouter()
	protR = protR.SkipClean(true)
	protR.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)
	protR.Use(middlewares.Auth(utils.NewJWTUtil(config.JWT.Secret)))

	// Authenticated routes
	protR.HandleFunc("/secure-health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.PathPrefix("").Handler(protR)
	protR.HandleFunc("/user/profile", userHandler.CreateUserProfile).Methods(http.MethodPost)
	protR.HandleFunc("/user/profile/address", userHandler.CreateUserAddress).Methods(http.MethodPost)
	protR.HandleFunc("/user/profile", userHandler.GetUserProfile).Methods(http.MethodGet)
	protR.HandleFunc("/user/change-password", userHandler.ChangePassword).Methods(http.MethodPut)
	protR.HandleFunc("/user/profile", userHandler.UpdateUserProfile).Methods(http.MethodPut)
	protR.HandleFunc("/user/email", userHandler.UpdateUserEmail).Methods(http.MethodPut)
	// ...
	// todo: should this be like get all future ones, or ones in the attendee's city?
	protR.HandleFunc("/festival", festivalHandler.GetAll).Methods(http.MethodGet)
	// ... ORGANIZER ONLY
	protR.HandleFunc("/festival", festivalHandler.Create).Methods(http.MethodPost)
	protR.HandleFunc("/organizer/festival", festivalHandler.GetByOrganizer).Methods(http.MethodGet)
	protR.HandleFunc("/festival/{festivalId}", festivalHandler.GetById).Methods(http.MethodGet)
	protR.HandleFunc("/festival/{festivalId}", festivalHandler.Update).Methods(http.MethodPut)
	protR.HandleFunc("/festival/{festivalId}", festivalHandler.Delete).Methods(http.MethodDelete)

	protR.HandleFunc("/festival/{festivalId}/publish", festivalHandler.PublishFestival).Methods(http.MethodPut)
	protR.HandleFunc("/festival/{festivalId}/cancel", festivalHandler.CancelFestival).Methods(http.MethodPut)
	protR.HandleFunc("/festival/{festivalId}/complete", festivalHandler.CompleteFestival).Methods(http.MethodPut)
	protR.HandleFunc("/festival/{festivalId}/store/open", festivalHandler.OpenStore).Methods(http.MethodPut)
	protR.HandleFunc("/festival/{festivalId}/store/close", festivalHandler.CloseStore).Methods(http.MethodPut)

	protR.HandleFunc("/festival/{festivalId}/image", festivalHandler.GetImages).Methods(http.MethodGet)
	protR.HandleFunc("/festival/{festivalId}/image", festivalHandler.AddImage).Methods(http.MethodPost)

	protR.HandleFunc("/organizer/employee", userHandler.CreateEmployee).Methods(http.MethodPost)
	protR.HandleFunc("/organizer/festival/{festivalId}/employee", userHandler.GetFestivalEmployees).Methods(http.MethodGet)
	protR.HandleFunc("/organizer/festival/{festivalId}/employee/{employeeId}/employ", festivalHandler.Employ).Methods(http.MethodPut)
	protR.HandleFunc("/organizer/festival/{festivalId}/employee/count", festivalHandler.GetEmployeeCount).Methods(http.MethodGet)
	protR.HandleFunc("/organizer/festival/{festivalId}/employee/available", userHandler.GetEmployeesNotOnFestival).Methods(http.MethodGet)

	// ...

	return r
}
