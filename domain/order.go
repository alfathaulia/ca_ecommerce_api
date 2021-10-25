package domain

import (
	"context"
	"time"
)

type Order struct {
	UserID        User      `json:"user_id" validate:"required"`
	PayMethod     string    `json:"paymethod" validate:"required"`
	TaxPrice      float32   `json:"tax_price" validate:"required"`
	ShippingPrice float32   `json:"shipping_price" validate:"required"`
	TotalPrice    float32   `json:"total_price" validate:"required"`
	IsPaid        bool      `json:"is_paid" validate:"required"`
	IsDelivered   bool      `json:"is_delivered" validate:"required"`
	PaidAt        time.Time `json:"paid_at" validate:"required"`
	DeliveredAt   time.Time `json:"delivered_at" validate:"required"`
	CreatedAt     time.Time `json:"created_at" validate:"required"`
}

type OrderRepository interface {
	Fetch(ctx context.Context, cursor string, num int) ([]Order, string, error)
	GetByID(ctx context.Context, id int) (Order, error)
	Update(ctx context.Context, updateOrder *Order) error
	Store(ctx context.Context, createOrder *Order) error
	Delete(ctx context.Context, id int) error
}

type OrderUsecase interface {
	Fetch(ctx context.Context, cursor string, num int) ([]Order, string, error)
	GetByID(ctx context.Context, id int) (Order, error)
	Update(ctx context.Context, updateOrder *Order) error
	Store(ctx context.Context, createOrder *Order) error
	Delete(ctx context.Context, id int) error
}
