package domain

import "context"

type OrderItem struct {
	ID        int64   `json:"id" verified:"required"`
	ProductID Product `json:"product_id" verified:"required"`
	OrderID   Order   `json:"order_id" verified:"required"`
	Name      string  `json:"name" verified:"required"`
	Qty       int     `json:"qty" verified:"required"`
	Price     float32 `json:"price" verified:"required"`
	Image     string  `json:"image" verified:"required"`
}

type OrderItemRepository interface {
	Fetch(ctx context.Context, cursor string, id int64) ([]OrderItem, string, error)
	GetByID(ctx context.Context, id int64) (OrderItem, error)
}
