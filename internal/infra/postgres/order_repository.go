package postgres

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        string    `gorm:"column:order_id;primaryKey"`
	UserID    int64     `gorm:"column:user_ref;index"`
	Amount    float64   `gorm:"column:total_amount"`
	Status    string    `gorm:"column:order_status"`
	CreatedAt time.Time `gorm:"column:creatd_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:update_at;autoUpdateTime"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateOrder(order *Order) error {
	return r.DB.Create(order).Error
}
