package repository

import (
	"order-management-service/internal/domain"
	"order-management-service/internal/port/outbound"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (OrderEntity) TableName() string {
	return "orders"
}

type OrderEntity struct {
	ID        string    `gorm:"column:order_id;primaryKey"`
	UserID    int64     `gorm:"column:user_ref;index"`
	Amount    float64   `gorm:"column:total_amount"`
	Status    string    `gorm:"column:order_status"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type Repository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) outbound.OrderRepository {
	return &Repository{DB: db}
}

func (r *Repository) Create(order *domain.Order) error {
	order.ID = uuid.NewString()
	return r.DB.Create(toEntity(order)).Error
}

func (r *Repository) GetByID(id string) (*domain.Order, error) {
	var entity OrderEntity

	if err := r.DB.First(&entity, "order_id = ?", id).Error; err != nil {
		return nil, err
	}

	return toDomain(&entity), nil
}

func toEntity(d *domain.Order) *OrderEntity {
	return &OrderEntity{
		ID:        d.ID,
		UserID:    d.UserID,
		Amount:    d.Amount,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func toDomain(e *OrderEntity) *domain.Order {
	return &domain.Order{
		ID:        e.ID,
		UserID:    e.UserID,
		Amount:    e.Amount,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
