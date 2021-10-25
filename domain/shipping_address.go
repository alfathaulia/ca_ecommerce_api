package domain

import "context"

type ShippingAddress struct {
	ID            int64   `json:"id" verified:"required"`
	OderID        Order   `json:"order_id" verified:"required"`
	Address       string  `json:"string" verified:"required"`
	City          string  `json:"city" verified:"required"`
	PostalCode    int     `json:"postal_code" verified:"required"`
	Country       string  `json:"country" verified:"required"`
	ShippingPrice float32 `json:"shipping_price" verified:"required"`
}

type ShippingAddressRepository interface {
	Fetch(ctx context.Context, cursor string, id int) ([]ShippingAddress, string, error)
	GetByID(ctx context.Context, id int64) (ShippingAddress, error)
}
