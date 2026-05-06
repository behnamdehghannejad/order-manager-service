package handler

import (
	"encoding/json"
	"net/http"

	"order-management-service/internal/port/inbound"
)

type CreateOrderRequest struct {
	UserID int64   `json:"user_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

type OrderHandler struct {
	useCase inbound.OrderUseCase
}

func NewOrderHandler(orderUseCase inbound.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: orderUseCase}
}

func (h *OrderHandler) CreateOrder(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req CreateOrderRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.useCase.Create(req.UserID, req.Amount, req.Status)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	order, err := h.useCase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
