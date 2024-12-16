package router

import (
	"backend/internal/config"
	handlersCommon "backend/internal/handlers/common"
	handlersFestival "backend/internal/handlers/festival"
	handlersUser "backend/internal/handlers/user"
	"backend/internal/middlewares"
	reposCommon "backend/internal/repositories/common"
	reposFestival "backend/internal/repositories/festival"
	reposUser "backend/internal/repositories/user"
	servicesCommon "backend/internal/services/common"
	servicesFestival "backend/internal/services/festival"
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
	logRepo := reposUser.NewLogRepo(db)
	userRepo := reposUser.NewUserRepo(db)
	userProfileRepo := reposUser.NewUserProfileRepo(db)
	addressRepo := reposCommon.NewAddressRepo(db)
	cityRepo := reposCommon.NewCityRepo(db)
	countryRepo := reposCommon.NewCountryRepo(db)
	festivalRepo := reposFestival.NewFestivalRepo(db)
	imageRepo := reposCommon.NewImageRepo(db)
	itemRepo := reposFestival.NewItemRepo(db)
	orderRepo := reposFestival.NewOrderRepository(db)
	// ...

	// Init services
	logger := servicesUser.NewLogger(logRepo, userRepo)
	logService := servicesUser.NewLogService(logRepo)
	pdfGenerator := servicesCommon.NewPDFGenerator()
	emailService := servicesCommon.NewEmailService(config)
	locationService := servicesCommon.NewLocationService(addressRepo, cityRepo, countryRepo)
	userService := servicesUser.NewUserService(config, userRepo, userProfileRepo, locationService, imageRepo)
	festivalService := servicesFestival.NewFestivalService(config, festivalRepo, userRepo, locationService, imageRepo, orderRepo)
	itemService := servicesFestival.NewItemService(config, itemRepo, locationService, imageRepo)
	awsService := servicesCommon.NewAWSService(s3Client, s3Presign, config)
	orderService := servicesFestival.NewOrderService(orderRepo, itemRepo, festivalRepo, userService, imageRepo, locationService, pdfGenerator)
	// ...

	// Init handlers
	logHandler := handlersUser.NewLogHandler(logService)
	commonHandler := handlersCommon.NewHealthHandler(config)
	userHandler := handlersUser.NewUserHandler(logger, userService, locationService)
	festivalHandler := handlersFestival.NewFestivalHandler(logger, festivalService, locationService)
	itemHandler := handlersFestival.NewItemHandler(logger, itemService, festivalService)
	awsHandler := handlersCommon.NewAWSHandler(logger, awsService, festivalService)
	orderHandler := handlersFestival.NewOrderHandler(logger, orderService, userService, emailService)
	// ...

	r := mux.NewRouter()
	r = r.SkipClean(true) // todo: see what this does
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	// Unauthenticated routes
	r.HandleFunc("/health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/user/login", userHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/user/register-attendee", userHandler.RegisterAttendee).Methods(http.MethodPost)
	r.HandleFunc("/festival/count", festivalHandler.GetFestivalCount).Methods(http.MethodGet)
	r.HandleFunc("/festival/attendee/count", festivalHandler.GetAttendeeCount).Methods(http.MethodGet)

	// ...

	pR := mux.NewRouter()
	pR = pR.SkipClean(true)
	pR.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)
	pR.Use(middlewares.Auth(utils.NewJWTUtil(config.JWT.Secret)))

	// Authenticated routes
	pR.HandleFunc("/secure-health", commonHandler.HealthCheck).Methods(http.MethodGet)
	r.PathPrefix("").Handler(pR) // ??

	pR.HandleFunc("/user", userHandler.GetUsers).Methods(http.MethodGet)
	pR.HandleFunc("/user/profile", userHandler.CreateUserProfile).Methods(http.MethodPost)
	pR.HandleFunc("/user/profile/address", userHandler.CreateUserAddress).Methods(http.MethodPost)
	pR.HandleFunc("/user/profile", userHandler.GetUserProfile).Methods(http.MethodGet)
	pR.HandleFunc("/user/change-password", userHandler.ChangePassword).Methods(http.MethodPut)
	pR.HandleFunc("/user/profile", userHandler.UpdateUserProfile).Methods(http.MethodPut)
	pR.HandleFunc("/user/email", userHandler.UpdateUserEmail).Methods(http.MethodPut)
	pR.HandleFunc("/user/profile/address", userHandler.UpdateUserAddress).Methods(http.MethodPut)
	pR.HandleFunc("/user/profile/photo", userHandler.UpdateProfilePhoto).Methods(http.MethodPut)
	pR.HandleFunc("/user/{userId}", userHandler.GetUser).Methods(http.MethodGet)

	// ...

	pR.HandleFunc("/festival", festivalHandler.GetAll).Methods(http.MethodGet)
	// ... ORGANIZER ONLY
	pR.HandleFunc("/festival", festivalHandler.Create).Methods(http.MethodPost)
	pR.HandleFunc("/festival/attendee", festivalHandler.GetAll).Methods(http.MethodGet)          // get all by organizer
	pR.HandleFunc("/festival/organizer", festivalHandler.GetByOrganizer).Methods(http.MethodGet) // get all by organizer
	pR.HandleFunc("/festival/employee", festivalHandler.GetByEmployee).Methods(http.MethodGet)   // get all by employee
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
	pR.HandleFunc("/festival/{festivalId}/image/{imageId}", festivalHandler.RemoveImage).Methods(http.MethodDelete)

	pR.HandleFunc("/organizer/employee", userHandler.CreateEmployee).Methods(http.MethodPost)
	pR.HandleFunc("/organizer/employee", userHandler.UpdateStaffProfile).Methods(http.MethodPut)
	pR.HandleFunc("/organizer/employee/email", userHandler.UpdateStaffEmail).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/employee", userHandler.GetFestivalEmployees).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/employee/{employeeId}/employ", festivalHandler.Employ).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/employee/{employeeId}/fire", festivalHandler.Fire).Methods(http.MethodDelete)
	pR.HandleFunc("/festival/{festivalId}/employee/count", festivalHandler.GetEmployeeCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/employee/available", userHandler.GetEmployeesNotOnFestival).Methods(http.MethodGet)

	// * ...ITEMS

	pR.HandleFunc("/festival/{festivalId}/item", itemHandler.CreateItem).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/item/price", itemHandler.CreatePriceListItem).Methods(http.MethodPost)

	// * ... TICKET TYPES

	pR.HandleFunc("/festival/{festivalId}/item/ticket-type", itemHandler.GetCurrentTicketTypes).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/ticket-type/count", itemHandler.GetTicketTypesCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/ticket-type/{itemId}", itemHandler.GetTicketType).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/ticket-type/{itemId}", itemHandler.UpdateItem).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/item/ticket-type/{itemId}", itemHandler.DeleteTicketType).Methods(http.MethodDelete)

	// * ... PACKAGE ADDONS

	pR.HandleFunc("/festival/{festivalId}/item/package-addon", itemHandler.CreatePackageAddon).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/{category}/count", itemHandler.GetPackageAddonsCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/count", itemHandler.GetAllPackageAddonsCount).Methods(http.MethodGet)

	pR.HandleFunc("/festival/{festivalId}/item/package-addon/general", itemHandler.GetGeneralAddons).Methods(http.MethodGet)

	pR.HandleFunc("/festival/{festivalId}/item/package-addon/transport", itemHandler.GetTransportAddons).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/transport", itemHandler.CreateTransportPackageAddon).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/transport/countries", itemHandler.GetAvailableDepartureCountries).Methods(http.MethodGet)

	pR.HandleFunc("/festival/{festivalId}/item/package-addon/camp", itemHandler.GetCampAddons).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/camp", itemHandler.CreateCampPackageAddon).Methods(http.MethodPost)

	// ...

	pR.HandleFunc("/image/upload", awsHandler.GetPresignedURL).Methods(http.MethodPost)

	// ...

	pR.HandleFunc("/order/attendee", orderHandler.GetOrdersAttendee).Methods(http.MethodGet)
	pR.HandleFunc("/order/{orderId}", orderHandler.GetOrder).Methods(http.MethodGet)
	pR.HandleFunc("/order/{orderId}/shipping-label", orderHandler.GetShippingLabel).Methods(http.MethodGet)

	pR.HandleFunc("/festival/{festivalId}/order", orderHandler.GetOrdersEmployee).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/order/count", orderHandler.GetOrdersCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/order/ticket", orderHandler.CreateTicketOrder).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/order/package", orderHandler.CreatePackageOrder).Methods(http.MethodPost)

	// ...

	pR.HandleFunc("/bracelet", orderHandler.IssueBracelet).Methods(http.MethodPost)
	pR.HandleFunc("/bracelet/attendee", orderHandler.GetBraceletOrdersAttendee).Methods(http.MethodGet)
	pR.HandleFunc("/bracelet/{braceletId}/activate/help", orderHandler.GetHelpRequest).Methods(http.MethodGet)
	pR.HandleFunc("/bracelet/{braceletId}/activate/help", orderHandler.SendActivateBraceletHelpRequest).Methods(http.MethodPost)
	pR.HandleFunc("/bracelet/{braceletId}/activate/help/approve", orderHandler.ApproveHelpRequest).Methods(http.MethodPut)
	pR.HandleFunc("/bracelet/{braceletId}/activate/help/reject", orderHandler.RejectHelpRequest).Methods(http.MethodPut)

	pR.HandleFunc("/bracelet/{braceletId}/activate", orderHandler.ActivateBracelet).Methods(http.MethodPut)
	pR.HandleFunc("/bracelet/{braceletId}/top-up", orderHandler.TopUpBracelet).Methods(http.MethodPut)

	// ...

	pR.HandleFunc("/admin/organizer", userHandler.CreateOrganizer).Methods(http.MethodPost)
	pR.HandleFunc("/admin/admin", userHandler.CreateAdmin).Methods(http.MethodPost)

	// ...

	pR.HandleFunc("/log", logHandler.GetAllLogs).Methods(http.MethodGet)
	pR.HandleFunc("/log/{role}", logHandler.GetLogsByRole).Methods(http.MethodGet)

	return r
}
