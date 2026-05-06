package domain

import "time"

type Order struct {
	ID        string
	UserID    int64
	Amount    float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
