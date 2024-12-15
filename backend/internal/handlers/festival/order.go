package handlers

import (
	dtoFestival "backend/internal/dto/festival"
	models "backend/internal/models/festival"
	servicesCommon "backend/internal/services/common"
	servicesFestival "backend/internal/services/festival"
	servicesUser "backend/internal/services/user"
	"backend/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type OrderHandler interface {
	CreateTicketOrder(w http.ResponseWriter, r *http.Request)
	CreatePackageOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	GetOrdersAttendee(w http.ResponseWriter, r *http.Request)
	GetOrdersEmployee(w http.ResponseWriter, r *http.Request)
	GetOrdersCount(w http.ResponseWriter, r *http.Request)
	IssueBracelet(w http.ResponseWriter, r *http.Request)
	GetBraceletOrdersAttendee(w http.ResponseWriter, r *http.Request)
	ActivateBracelet(w http.ResponseWriter, r *http.Request)
	TopUpBracelet(w http.ResponseWriter, r *http.Request)
	SendActivateBraceletHelpRequest(w http.ResponseWriter, r *http.Request)
	GetHelpRequest(w http.ResponseWriter, r *http.Request)
	ApproveHelpRequest(w http.ResponseWriter, r *http.Request)
	RejectHelpRequest(w http.ResponseWriter, r *http.Request)
	GetShippingLabel(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	log          servicesCommon.Logger
	orderService servicesFestival.OrderService
	userService  servicesUser.UserService
	emailService servicesCommon.EmailService
}

func NewOrderHandler(
	lg servicesCommon.Logger,
	os servicesFestival.OrderService,
	us servicesUser.UserService,
	es servicesCommon.EmailService,
) OrderHandler {
	return &orderHandler{
		orderService: os,
		userService:  us,
		emailService: es,
		log:          lg,
	}
}

func (h *orderHandler) CreateTicketOrder(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoFestival.CreateTicketOrderRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	festivalTicket := models.FestivalTicket{
		ItemID: input.TicketTypeId,
	}
	if err := h.orderService.CreateFestivalTicket(&festivalTicket); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	username := utils.GetUsername(r.Context())
	attendeeId, err := h.userService.GetUserID(username)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	order := models.Order{
		FestivalTicketID: festivalTicket.ID,
		TotalAmount:      float64(input.TotalPrice),
		UserID:           attendeeId,
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	email := h.userService.GetUserEmail(username)
	if email != "" {
		if err := h.emailService.SendEmail(email, "Order Created", fmt.Sprintf("Ticket Order #%d created successfully!", order.ID)); err != nil {
			log.Println("error:", err)
		}
	} else {
		log.Println("email not found for user", username)
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"orderId": order.ID}, nil)
	h.log.Info("ticket order created: "+fmt.Sprint(order.ID), r.Context())
}

func (h *orderHandler) CreatePackageOrder(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoFestival.CreatePackageOrderRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	festivalTicket := models.FestivalTicket{
		ItemID: input.TicketTypeId,
	}
	if err := h.orderService.CreateFestivalTicket(&festivalTicket); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	festivalPackage := models.FestivalPackage{}
	if err := h.orderService.CreateFestivalPackage(&festivalPackage); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// add transport addon
	if input.TransportAddonId != nil && *input.TransportAddonId != 0 {
		festivalPackageAddonTransport := models.FestivalPackageAddon{
			ItemID:            *input.TransportAddonId,
			FestivalPackageID: festivalPackage.ID,
		}
		if err := h.orderService.CreateFestivalPackageAddon(&festivalPackageAddonTransport); err != nil {
			log.Println("error:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	// add camp addon
	if input.CampAddonId != nil && *input.CampAddonId != 0 {
		festivalPackageAddonCamp := models.FestivalPackageAddon{
			ItemID:            *input.CampAddonId,
			FestivalPackageID: festivalPackage.ID,
		}
		if err := h.orderService.CreateFestivalPackageAddon(&festivalPackageAddonCamp); err != nil {
			log.Println("error:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	// add general addons
	if input.GeneralAddonIds != nil && len(*input.GeneralAddonIds) > 0 {
		for _, addon := range *input.GeneralAddonIds {
			festivalPackageAddon := models.FestivalPackageAddon{
				ItemID:            addon,
				FestivalPackageID: festivalPackage.ID,
			}
			if err := h.orderService.CreateFestivalPackageAddon(&festivalPackageAddon); err != nil {
				log.Println("error:", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
	}

	username := utils.GetUsername(r.Context())
	attendeeId, err := h.userService.GetUserID(username)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	order := models.Order{
		FestivalTicketID:  festivalTicket.ID,
		FestivalPackageID: &festivalPackage.ID,
		TotalAmount:       float64(input.TotalPrice),
		UserID:            attendeeId,
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	email := h.userService.GetUserEmail(username)
	if email != "" {
		if err := h.emailService.SendEmail(email, "Order Created", fmt.Sprintf("Package Order #%d created successfully!", order.ID)); err != nil {
			log.Println("error:", err)
		}
	} else {
		log.Println("email not found for user", username)
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"orderId": order.ID}, nil)
	h.log.Info("package order created: "+fmt.Sprint(order.ID), r.Context())
}

func (h *orderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {

	isAttendee := utils.AuthAttendeeRole(r.Context())
	isEmployeeOrOrganizer := utils.AuthEmployeeRole(r.Context()) || utils.AuthOrganizerRole(r.Context())

	if !isAttendee && !isEmployeeOrOrganizer {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var username string
	if isAttendee {
		username = utils.GetUsername(r.Context())
	} else if isEmployeeOrOrganizer {
		username = ""
	}

	orderId, err := GetIDParamFromRequest(r, "orderId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	order, err := h.orderService.GetOrder(username, orderId)
	if err != nil {
		log.Println("error:", err)
		if err.Error() == "order not found" || err.Error() == "record not found" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, order, nil)
	log.Println("order fetched", order.OrderID)
}

func (h *orderHandler) GetOrdersAttendee(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	username := utils.GetUsername(r.Context())

	orders, err := h.orderService.GetOrdersAttendee(username)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders, nil)
	log.Println("orders fetched for user", username)
}

func (h *orderHandler) GetOrdersEmployee(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthEmployeeRole(r.Context()) && !utils.AuthOrganizerRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	festivalId, err := GetIDParamFromRequest(r, "festivalId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	orders, err := h.orderService.GetOrdersEmployee(festivalId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders, nil)
	log.Println("orders fetched for festival", festivalId)
}

func (h *orderHandler) GetOrdersCount(w http.ResponseWriter, r *http.Request) {

	if !utils.Auth(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	festivalId, err := GetIDParamFromRequest(r, "festivalId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	count, err := h.orderService.GetOrdersCount(festivalId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, dtoFestival.FestivalPropCountResponse{
		FestivalId: festivalId,
		Count:      int(count),
	}, nil)
	log.Println("order count retrieved successfully for festival:", festivalId)
}

func (h *orderHandler) IssueBracelet(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthEmployeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var input dtoFestival.IssueBraceletRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	employeeId, err := h.userService.GetUserID(utils.GetUsername(r.Context()))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	attendee, err := h.userService.GetUserProfile(input.AttendeeUsername)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	attendeeId, err := h.userService.GetUserID(input.AttendeeUsername)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	bracelet := models.Bracelet{
		PIN:              input.PIN,
		BarcodeNumber:    input.BarcodeNumber,
		Balance:          0,
		Status:           models.BraceletStatusIssued,
		FestivalTicketID: input.FestivalTicketId,
		AttendeeID:       attendeeId,
		EmployeeID:       employeeId,
	}

	if err := h.orderService.IssueBracelet(&bracelet); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	email := h.userService.GetUserEmail(attendee.Username)
	if email != "" {
		if err := h.emailService.SendEmail(
			email,
			"Bracelet Issued",
			fmt.Sprintf("Bracelet for Order #%d issued and sent to your address!", input.OrderId),
		); err != nil {
			log.Println("error:", err)
		}
	} else {
		log.Println("email not found for user", attendee.Username)
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"braceletId": bracelet.ID, "shippingAddress": attendee.Address}, nil)
	h.log.Info("bracelet issued: "+fmt.Sprint(bracelet.ID), r.Context())
}

func (h *orderHandler) GetBraceletOrdersAttendee(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	username := utils.GetUsername(r.Context())

	orders, err := h.orderService.GetBraceletOrdersAttendee(username)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders, nil)
	log.Println("orders fetched for user", username)
}

func (h *orderHandler) ActivateBracelet(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	username := utils.GetUsername(r.Context())

	braceletId, err := GetIDParamFromRequest(r, "braceletId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var input dtoFestival.ActivateBraceletRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.orderService.ActivateBracelet(username, braceletId, input.PIN); err != nil {
		log.Println("error:", err)
		if strings.Contains(err.Error(), "bracelet not found") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else if strings.Contains(err.Error(), "invalid PIN") {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	email := h.userService.GetUserEmail(username)
	if email != "" {
		if err := h.emailService.SendEmail(
			email,
			"Bracelet Activated",
			"Bracelet activated successfully!",
		); err != nil {
			log.Println("error:", err)
		}
	} else {
		log.Println("email not found for user", username)
	}

	utils.WriteJSON(w, http.StatusOK, nil, nil)
	h.log.Info("bracelet activated: "+fmt.Sprint(braceletId), r.Context())
}

func (h *orderHandler) TopUpBracelet(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	username := utils.GetUsername(r.Context())

	braceletId, err := GetIDParamFromRequest(r, "braceletId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var input dtoFestival.TopUpBraceletRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.orderService.TopUpBracelet(username, braceletId, input.Amount); err != nil {
		log.Println("error:", err)
		if strings.Contains(err.Error(), "bracelet not found") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else if strings.Contains(err.Error(), "invalid PIN") {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	email := h.userService.GetUserEmail(username)
	if email != "" {
		if err := h.emailService.SendEmail(
			email,
			"Bracelet Top Up Receipt",
			fmt.Sprintf("Payment for Bracelet top up is successful! Top Up amount: $%.2f", input.Amount),
		); err != nil {
			log.Println("error:", err)
		}
	} else {
		log.Println("email not found for user", username)
	}

	utils.WriteJSON(w, http.StatusOK, nil, nil)
	h.log.Info("bracelet top-up: "+fmt.Sprint(braceletId), r.Context())
}

func (h *orderHandler) SendActivateBraceletHelpRequest(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthAttendeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	username := utils.GetUsername(r.Context())

	braceletId, err := GetIDParamFromRequest(r, "braceletId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var input dtoFestival.ActivateBraceletHelpRequest
	if err := utils.ReadJSON(w, r, &input); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.orderService.CreateHelpRequest(username, input); err != nil {
		log.Println("error:", err)
		if strings.Contains(err.Error(), "bracelet not found") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	utils.WriteJSON(w, http.StatusOK, nil, nil)
	h.log.Info("bracelet activation help requested: "+fmt.Sprint(braceletId), r.Context())
}

func (h *orderHandler) GetHelpRequest(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthEmployeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	braceletId, err := GetIDParamFromRequest(r, "braceletId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	helpRequest, err := h.orderService.GetHelpRequest(braceletId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	employeeBracelet, err := h.userService.GetUserProfileById(helpRequest.Bracelet.EmployeeID)
	if err != nil {
		log.Println("error: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	attendee, err := h.userService.GetUserProfileById(helpRequest.Bracelet.AttendeeID)
	if err != nil {
		log.Println("error: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	bracelet := dtoFestival.BraceletDTO{
		BraceletID:    helpRequest.Bracelet.ID,
		BarcodeNumber: helpRequest.Bracelet.BarcodeNumber,
		Status:        helpRequest.Bracelet.Status,
		Balance:       helpRequest.Bracelet.Balance,
		Employee:      employeeBracelet,
		PIN:           &helpRequest.Bracelet.PIN,
	}

	response := dtoFestival.ActivationHelpRequestDTO{
		ActivationHelpRequestID: helpRequest.ID,
		UserEnteredPIN:          helpRequest.UserEnteredPIN,
		UserEnteredBarcode:      helpRequest.UserEnteredBarcode,
		IssueDescription:        helpRequest.IssueDescription,
		ImageURL:                helpRequest.ProofImage.URL,
		Status:                  helpRequest.Status,
		Bracelet:                bracelet,
		Attendee:                *attendee,
	}

	utils.WriteJSON(w, http.StatusOK, response, nil)
	log.Println("help request fetched for bracelet", braceletId)
}

func (h *orderHandler) ApproveHelpRequest(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthEmployeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	braceletId, err := GetIDParamFromRequest(r, "braceletId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.orderService.ApproveHelpRequest(braceletId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	bracelet, err := h.orderService.GetBraceletById(braceletId)
	if err == nil {
		if err := h.emailService.SendEmail(
			bracelet.Attendee.User.Email,
			"Bracelet Activated (via Help Request)",
			"Employee approved your activation help request and Bracelet is now activated!",
		); err != nil {
			log.Println("error:", err)
		}
	} else {
		log.Println("error:", err)
	}

	utils.WriteJSON(w, http.StatusOK, nil, nil)
	h.log.Info("bracelet activation help request approved: "+fmt.Sprint(bracelet.ID), r.Context())
}

func (h *orderHandler) RejectHelpRequest(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthEmployeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	braceletId, err := GetIDParamFromRequest(r, "braceletId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.orderService.RejectHelpRequest(braceletId); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	bracelet, err := h.orderService.GetBraceletById(braceletId)
	if err == nil {
		if err := h.emailService.SendEmail(
			bracelet.Attendee.User.Email,
			"Bracelet Activation Rejected (via Help Request)",
			"Employee rejected your activation help request and Bracelet is now in rejected state.",
		); err != nil {
			log.Println("error:", err)
		}
	} else {
		log.Println("error:", err)
	}

	utils.WriteJSON(w, http.StatusOK, nil, nil)
	h.log.Info("bracelet activation help request rejected: "+fmt.Sprint(bracelet.ID), r.Context())
}

func (h *orderHandler) GetShippingLabel(w http.ResponseWriter, r *http.Request) {

	if !utils.AuthEmployeeRole(r.Context()) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	orderId, err := GetIDParamFromRequest(r, "orderId")
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	pdfBytes, err := h.orderService.GetShippingLabel(orderId)
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=shipping_label.pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
	log.Println("shipping label fetched for order", orderId)
}
