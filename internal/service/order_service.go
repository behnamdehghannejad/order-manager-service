package service

import (
	"order-management-service/internal/domain"
	"order-management-service/internal/port/inbound"
	"order-management-service/internal/port/outbound"
	"order-management-service/internal/utility"

	"github.com/google/uuid"
)

type OrderService struct {
	repo outbound.OrderRepository
}

func NewOrderService(orderRepository outbound.OrderRepository) inbound.OrderUseCase {
	return &OrderService{repo: orderRepository}
}

func (s *OrderService) Create(userID int64, amount float64, status string) (*domain.Order, error) {

	if err := utility.ValidateOrderInput(userID, amount, status); err != nil {
		return nil, err
	}

	order := &domain.Order{
		ID:     uuid.NewString(),
		UserID: userID,
		Amount: amount,
		Status: status,
	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetByID(id string) (*domain.Order, error) {
	return s.repo.GetByID(id)
}
