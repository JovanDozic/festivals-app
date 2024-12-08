package handlers

import (
	dtoFestival "backend/internal/dto/festival"
	models "backend/internal/models/festival"
	servicesFestival "backend/internal/services/festival"
	servicesUser "backend/internal/services/user"
	"backend/internal/utils"
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
	IssueBracelet(w http.ResponseWriter, r *http.Request)
	GetBraceletOrdersAttendee(w http.ResponseWriter, r *http.Request)
	ActivateBracelet(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	orderService servicesFestival.OrderService
	userService  servicesUser.UserService
}

func NewOrderHandler(os servicesFestival.OrderService, us servicesUser.UserService) OrderHandler {
	return &orderHandler{orderService: os, userService: us}
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

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"orderId": order.ID}, nil)
	log.Println("order created", order.ID)
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

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"orderId": order.ID}, nil)
	log.Println("order created", order.ID)
}

func (h *orderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {

	isAttendee := utils.AuthAttendeeRole(r.Context())
	isEmployee := utils.AuthEmployeeRole(r.Context())

	if !isAttendee && !isEmployee {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var username string
	if isAttendee {
		username = utils.GetUsername(r.Context())
	} else if isEmployee {
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

	if !utils.AuthEmployeeRole(r.Context()) {
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

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"braceletId": bracelet.ID, "shippingAddress": attendee.Address}, nil)
	log.Println("bracelet issued", bracelet.ID)
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

	utils.WriteJSON(w, http.StatusOK, nil, nil)
	log.Println("bracelet activated", braceletId, "for user", username)
}
