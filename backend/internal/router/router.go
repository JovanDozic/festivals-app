package router

import (
	"backend/internal/container"
	"backend/internal/middlewares"
	"backend/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Init(c *container.Container) *mux.Router {

	// ! Unauthenticated routes

	r := mux.NewRouter()
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	r.HandleFunc("/user/login", c.Handlers.User.Login).Methods(http.MethodPost)
	r.HandleFunc("/user/register-attendee", c.Handlers.User.RegisterAttendee).Methods(http.MethodPost)
	r.HandleFunc("/festival/count", c.Handlers.Festival.GetFestivalCount).Methods(http.MethodGet)
	r.HandleFunc("/festival/attendee/count", c.Handlers.Festival.GetAttendeeCount).Methods(http.MethodGet)

	// ! Authenticated routes

	pR := mux.NewRouter()
	pR.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)
	pR.Use(middlewares.Auth(utils.NewJWTUtil(c.Config.JWT.Secret)))
	r.PathPrefix("").Handler(pR)

	// * ...USER

	pR.HandleFunc("/user", c.Handlers.User.GetUsers).Methods(http.MethodGet)
	pR.HandleFunc("/user/profile", c.Handlers.User.CreateUserProfile).Methods(http.MethodPost)
	pR.HandleFunc("/user/profile/address", c.Handlers.User.CreateUserAddress).Methods(http.MethodPost)
	pR.HandleFunc("/user/profile", c.Handlers.User.GetUserProfile).Methods(http.MethodGet)
	pR.HandleFunc("/user/change-password", c.Handlers.User.ChangePassword).Methods(http.MethodPut)
	pR.HandleFunc("/user/profile", c.Handlers.User.UpdateUserProfile).Methods(http.MethodPut)
	pR.HandleFunc("/user/email", c.Handlers.User.UpdateUserEmail).Methods(http.MethodPut)
	pR.HandleFunc("/user/profile/address", c.Handlers.User.UpdateUserAddress).Methods(http.MethodPut)
	pR.HandleFunc("/user/profile/photo", c.Handlers.User.UpdateProfilePhoto).Methods(http.MethodPut)
	pR.HandleFunc("/user/{userId}", c.Handlers.User.GetUser).Methods(http.MethodGet)

	// * ...IMAGES

	pR.HandleFunc("/image/upload", c.Handlers.AWS.GetPresignedURL).Methods(http.MethodPost)

	// * ... FESTIVALS

	// ** ... FESTIVALS

	pR.HandleFunc("/festival", c.Handlers.Festival.GetAll).Methods(http.MethodGet)
	pR.HandleFunc("/festival", c.Handlers.Festival.Create).Methods(http.MethodPost)
	pR.HandleFunc("/festival/attendee", c.Handlers.Festival.GetAll).Methods(http.MethodGet)
	pR.HandleFunc("/festival/organizer", c.Handlers.Festival.GetByOrganizer).Methods(http.MethodGet)
	pR.HandleFunc("/festival/organizer/{organizerId}", c.Handlers.Festival.GetByOrganizerById).Methods(http.MethodGet)
	pR.HandleFunc("/festival/employee", c.Handlers.Festival.GetByEmployee).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}", c.Handlers.Festival.GetById).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}", c.Handlers.Festival.Update).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}", c.Handlers.Festival.Delete).Methods(http.MethodDelete)

	// ** ... FESTIVAL STATUSES

	pR.HandleFunc("/festival/{festivalId}/publish", c.Handlers.Festival.PublishFestival).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/cancel", c.Handlers.Festival.CancelFestival).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/complete", c.Handlers.Festival.CompleteFestival).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/store/open", c.Handlers.Festival.OpenStore).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/store/close", c.Handlers.Festival.CloseStore).Methods(http.MethodPut)

	// ** ... FESTIVAL IMAGES

	pR.HandleFunc("/festival/{festivalId}/image", c.Handlers.Festival.GetImages).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/image", c.Handlers.Festival.AddImage).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/image/{imageId}", c.Handlers.Festival.RemoveImage).Methods(http.MethodDelete)

	// ** ... FESTIVAL EMPLOYEES

	pR.HandleFunc("/organizer/employee", c.Handlers.User.CreateEmployee).Methods(http.MethodPost)
	pR.HandleFunc("/organizer/employee", c.Handlers.User.UpdateStaffProfile).Methods(http.MethodPut)
	pR.HandleFunc("/organizer/employee/email", c.Handlers.User.UpdateStaffEmail).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/employee", c.Handlers.User.GetFestivalEmployees).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/employee/{employeeId}/employ", c.Handlers.Festival.Employ).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/employee/{employeeId}/fire", c.Handlers.Festival.Fire).Methods(http.MethodDelete)
	pR.HandleFunc("/festival/{festivalId}/employee/count", c.Handlers.Festival.GetEmployeeCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/employee/available", c.Handlers.User.GetEmployeesNotOnFestival).Methods(http.MethodGet)

	// ** ... FESTIVAL ITEMS

	pR.HandleFunc("/festival/{festivalId}/item", c.Handlers.Item.CreateItem).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/item/price", c.Handlers.Item.CreatePriceListItem).Methods(http.MethodPost)

	// ** ... TICKET TYPES

	pR.HandleFunc("/festival/{festivalId}/item/ticket-type", c.Handlers.Item.GetCurrentTicketTypes).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/ticket-type/count", c.Handlers.Item.GetTicketTypesCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/ticket-type/{itemId}", c.Handlers.Item.GetTicketType).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/ticket-type/{itemId}", c.Handlers.Item.UpdateItem).Methods(http.MethodPut)
	pR.HandleFunc("/festival/{festivalId}/item/ticket-type/{itemId}", c.Handlers.Item.DeleteTicketType).Methods(http.MethodDelete)

	// ** ... PACKAGE ADDONS

	pR.HandleFunc("/festival/{festivalId}/item/package-addon", c.Handlers.Item.CreatePackageAddon).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/{category}/count", c.Handlers.Item.GetPackageAddonsCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/count", c.Handlers.Item.GetAllPackageAddonsCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/general", c.Handlers.Item.GetGeneralAddons).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/transport", c.Handlers.Item.GetTransportAddons).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/transport", c.Handlers.Item.CreateTransportPackageAddon).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/transport/countries", c.Handlers.Item.GetAvailableDepartureCountries).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/camp", c.Handlers.Item.GetCampAddons).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/item/package-addon/camp", c.Handlers.Item.CreateCampPackageAddon).Methods(http.MethodPost)

	// * ... ORDERS

	pR.HandleFunc("/order/attendee", c.Handlers.Order.GetOrdersAttendee).Methods(http.MethodGet)
	pR.HandleFunc("/order/{orderId}", c.Handlers.Order.GetOrder).Methods(http.MethodGet)
	pR.HandleFunc("/order/{orderId}/shipping-label", c.Handlers.Order.GetShippingLabel).Methods(http.MethodGet)

	pR.HandleFunc("/festival/{festivalId}/order", c.Handlers.Order.GetOrdersEmployee).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/order/count", c.Handlers.Order.GetOrdersCount).Methods(http.MethodGet)
	pR.HandleFunc("/festival/{festivalId}/order/ticket", c.Handlers.Order.CreateTicketOrder).Methods(http.MethodPost)
	pR.HandleFunc("/festival/{festivalId}/order/package", c.Handlers.Order.CreatePackageOrder).Methods(http.MethodPost)

	// * ... BRACELETS

	pR.HandleFunc("/bracelet", c.Handlers.Order.IssueBracelet).Methods(http.MethodPost)
	pR.HandleFunc("/bracelet/attendee", c.Handlers.Order.GetBraceletOrdersAttendee).Methods(http.MethodGet)
	pR.HandleFunc("/bracelet/{braceletId}/activate/help", c.Handlers.Order.GetHelpRequest).Methods(http.MethodGet)
	pR.HandleFunc("/bracelet/{braceletId}/activate/help", c.Handlers.Order.SendActivateBraceletHelpRequest).Methods(http.MethodPost)
	pR.HandleFunc("/bracelet/{braceletId}/activate/help/approve", c.Handlers.Order.ApproveHelpRequest).Methods(http.MethodPut)
	pR.HandleFunc("/bracelet/{braceletId}/activate/help/reject", c.Handlers.Order.RejectHelpRequest).Methods(http.MethodPut)
	pR.HandleFunc("/bracelet/{braceletId}/activate", c.Handlers.Order.ActivateBracelet).Methods(http.MethodPut)
	pR.HandleFunc("/bracelet/{braceletId}/top-up", c.Handlers.Order.TopUpBracelet).Methods(http.MethodPut)

	// * ... ADMIN

	pR.HandleFunc("/admin/organizer", c.Handlers.User.CreateOrganizer).Methods(http.MethodPost)
	pR.HandleFunc("/admin/admin", c.Handlers.User.CreateAdmin).Methods(http.MethodPost)

	// * ... LOGS

	pR.HandleFunc("/log", c.Handlers.Log.GetAllLogs).Methods(http.MethodGet)
	pR.HandleFunc("/log/{role}", c.Handlers.Log.GetLogsByRole).Methods(http.MethodGet)

	return r
}
