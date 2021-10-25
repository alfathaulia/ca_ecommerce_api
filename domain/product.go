package domain

import (
	"context"
	"time"
)

type Product struct {
	ID           int64     `json:"id" `
	UserID       User      `json:"user_id"`
	Image        string    `json:"image" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Brand        string    `json:"brand" validate:"required"`
	Category     string    `json:"category" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Rating       int       `json:"rating" validate:"required"`
	NumReviews   int       `json:"num_reviews" validate:"required"`
	Price        int       `json:"price" validate:"required"`
	CountInStock int       `json:"count_in_stock" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
}

type ProductUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Product, string, error)
	GetByID(ctx context.Context, id int64) (Product, error)
	Update(ctx context.Context, ar *Product) error
	Store(context.Context, *Product) error
	Delete(ctx context.Context, id int64) error
}

type ProductRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Product, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Product, error)
	Update(ctx context.Context, ar *Product) error
	Store(ctx context.Context, a *Product) error
	Delete(ctx context.Context, id int64) error
}
