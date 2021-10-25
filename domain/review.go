package domain

import (
	"context"
	"time"
)

type Review struct {
	ID        int64     `json:"id" verified:"required"`
	ProductID Product   `json:"product_id" verified:"required"`
	UserID    User      `json:"user_id" verified:"required"`
	Name      string    `json:"name" verified:"required"`
	Rating    int       `json:"rating" verified:"required"`
	Comment   string    `json:"comment" verified:"required"`
	CreatedAt time.Time `json:"created_at" verified:"required"`
}

type ReviewRepository interface {
	Fetch(ctx context.Context, cursor string, num int) ([]Review, string, error)
	GetByID(ctx context.Context, id int64) (Review, error)
}
