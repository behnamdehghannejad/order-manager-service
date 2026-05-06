package inbound

import "order-management-service/internal/domain"

type OrderUseCase interface {
	Create(userID int64, amount float64, status string) (*domain.Order, error)
	GetByID(id string) (*domain.Order, error)
}
