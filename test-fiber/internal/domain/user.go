package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id" gorm:"primaryKey"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	Create(user *User) error
}

type UserUsecase interface {
	GetAll() ([]User, error)
	Create(request *CreateUserRequest) (*User, error)
}