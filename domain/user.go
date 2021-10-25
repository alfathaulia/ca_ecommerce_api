package domain

import (
	"context"
	"time"
)

type RolesType string

const (
	RolesTypeAdmin      RolesType = "admin"
	RolesTypeUser       RolesType = "user"
	RolesTypeStaff      RolesType = "staff"
	RolesTypeSuperadmin RolesType = "superadmin"
	RolesTypeSuperstaff RolesType = "superstaff"
)

// User ...
type User struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username" validate:"required"`
	Email          string    `json:"email" validate:"required"`
	HashedPassword string    `json:"hashed_password" validate:"required"`
	IsVerified     bool      `json:"is_verified" validate:"required"`
	Role           string    `json:"role"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// UserUsecase represent the User's usecases
type UserUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]User, string, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, ar *User) error
	GetByUsername(ctx context.Context, username string) (User, error)
	Store(context.Context, *User) error
	Delete(ctx context.Context, id int64) error
	Register(ctx context.Context, users *User) (err error)
	Login(ctx context.Context, username string, password string) (User, error)
	CreateAdmin(ctx context.Context, user *User) (err error)
	CreateStaff(ctx context.Context, user *User) (err error)
}

// UserRepository represent the User's repository contract
type UserRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []User, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	Update(ctx context.Context, ar *User) error
	Store(ctx context.Context, a *User) error
	Delete(ctx context.Context, id int64) error
	Register(ctx context.Context, users *User) (err error)
	Login(ctx context.Context, username string, password string) (User, error)
}
