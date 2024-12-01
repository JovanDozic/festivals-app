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
	"context"
	"log"
	"net/http"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, config *config.Config) *mux.Router {

	// Init AWS
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(config.AWS.Region),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(config.AWS.AccessKeyID, config.AWS.SecretAccessKey, ""),
		),
	)
	if err != nil {
		log.Fatalf("error initializing AWS config: %v", err)
		panic(err)
	}

	s3Client := s3.NewFromConfig(awsCfg)
	s3Presign := s3.NewPresignClient(s3Client)

	// Init repositories
	userRepo := reposUser.NewUserRepo(db)
	userProfileRepo := reposUser.NewUserProfileRepo(db)
	addressRepo := reposCommon.NewAddressRepo(db)
	cityRepo := reposCommon.NewCityRepo(db)
	countryRepo := reposCommon.NewCountryRepo(db)
	festivalRepo := reposFestival.NewFestivalRepo(db)
	imageRepo := reposCommon.NewImageRepo(db)
	itemRepo := reposFestival.NewItemRepo(db)
	// ...

	// Init services
	locationService := servicesCommon.NewLocationService(addressRepo, cityRepo, countryRepo)
	userService := servicesUser.NewUserService(config, userRepo, userProfileRepo, locationService, imageRepo)
	festivalService := services.NewFestivalService(config, festivalRepo, userRepo, locationService, imageRepo)
	itemService := services.NewItemService(config, itemRepo)
	awsService := servicesCommon.NewAWSService(s3Client, s3Presign, config)
	// ...

	// Init handlers
	commonHandler := handlersCommon.NewHealthHandler(config)
	userHandler := handlersUser.NewUserHandler(userService)
	festivalHandler := handlers.NewFestivalHandler(festivalService, locationService)
	itemHandler := handlers.NewItemHandler(itemService, festivalService)
	awsHandler := handlersCommon.NewAWSHandler(awsService, festivalService)
	// ...

	r := mux.NewRouter()
	r = r.SkipClean(true) // todo: see what this does
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	// Unauthenticated routes
	r.HandleFunc("/health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/user/login", userHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/user/register-attendee", userHandler.RegisterAttendee).Methods(http.MethodPost)
	// ...

	pR := mux.NewRouter()
	pR = pR.SkipClean(true)
	pR.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)
	pR.Use(middlewares.Auth(utils.NewJWTUtil(config.JWT.Secret)))

	// Authenticated routes
	pR.HandleFunc("/secure-health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.PathPrefix("").Handler(pR)
	pR.HandleFunc("/user/profile", userHandler.CreateUserProfile).Methods(http.MethodPost)
	pR.HandleFunc("/user/profile/address", userHandler.CreateUserAddress).Methods(http.MethodPost)
	pR.HandleFunc("/user/profile", userHandler.GetUserProfile).Methods(http.MethodGet)
	pR.HandleFunc("/user/change-password", userHandler.ChangePassword).Methods(http.MethodPut)
	pR.HandleFunc("/user/profile", userHandler.UpdateUserProfile).Methods(http.MethodPut)
	pR.HandleFunc("/user/email", userHandler.UpdateUserEmail).Methods(http.MethodPut)
	pR.HandleFunc("/user/profile/photo", userHandler.UpdateProfilePhoto).Methods(http.MethodPut)
	// ...
	// todo: should this be like get all future ones, or ones in the attendee's city?
	pR.HandleFunc("/festival", festivalHandler.GetAll).Methods(http.MethodGet)
	// ... ORGANIZER ONLY
	pR.HandleFunc("/festival", festivalHandler.Create).Methods(http.MethodPost)
	pR.HandleFunc("/organizer/festival", festivalHandler.GetByOrganizer).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}", festivalHandler.GetById).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}", festivalHandler.Update).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}", festivalHandler.Delete).Methods(http.MethodDelete)

	pR.HandleFunc("/festival/{festivalId}/publish", festivalHandler.PublishFestival).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/cancel", festivalHandler.CancelFestival).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/complete", festivalHandler.CompleteFestival).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/store/open", festivalHandler.OpenStore).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/store/close", festivalHandler.CloseStore).Methods(http.MethodPut)

	pR.HandleFunc("/festival/{festivalId}/image", festivalHandler.GetImages).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/image", festivalHandler.AddImage).Methods(http.MethodPost)

	pR.HandleFunc("/organizer/employee", userHandler.CreateEmployee).Methods(http.MethodPost)
	pR.HandleFunc("/organizer/employee", userHandler.UpdateStaffProfile).Methods(http.MethodPut)
	pR.HandleFunc("/organizer/employee/email", userHandler.UpdateStaffEmail).Methods(http.MethodPut)
	pR.HandleFunc("/organizer/festival/{festivalId}/employee", userHandler.GetFestivalEmployees).Methods(http.MethodGet)
	pR.HandleFunc("/organizer/festival/{festivalId}/employee/{employeeId}/employ", festivalHandler.Employ).Methods(http.MethodPut)
	pR.HandleFunc("/organizer/festival/{festivalId}/employee/{employeeId}/fire", festivalHandler.Fire).Methods(http.MethodDelete)
	pR.HandleFunc("/organizer/festival/{festivalId}/employee/count", festivalHandler.GetEmployeeCount).Methods(http.MethodGet)
	pR.HandleFunc("/organizer/festival/{festivalId}/employee/available", userHandler.GetEmployeesNotOnFestival).Methods(http.MethodGet)

	// ...

	pR.HandleFunc("/organizer/festival/{festivalId}/item", itemHandler.CreateItem).Methods(http.MethodPost)
	pR.HandleFunc("/organizer/festival/{festivalId}/item/price", itemHandler.CreatePriceListItem).Methods(http.MethodPost)
	pR.HandleFunc("/organizer/festival/{festivalId}/item/ticket-type", itemHandler.GetCurrentTicketTypes).Methods(http.MethodGet)
	pR.HandleFunc("/organizer/festival/{festivalId}/item/ticket-type/count", itemHandler.GetTicketTypesCount).Methods(http.MethodGet)
	pR.HandleFunc("/organizer/festival/{festivalId}/item/ticket-type/{itemId}", itemHandler.GetTicketType).Methods(http.MethodGet)
	pR.HandleFunc("/organizer/festival/{festivalId}/item/ticket-type/{itemId}", itemHandler.UpdateItem).Methods(http.MethodPut)
	pR.HandleFunc("/organizer/festival/{festivalId}/item/ticket-type/{itemId}", itemHandler.DeleteTicketType).Methods(http.MethodDelete)

	// ...

	pR.HandleFunc("/image/upload", awsHandler.GetPresignedURL).Methods(http.MethodPost)

	return r
}
