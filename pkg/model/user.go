package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `validate:"required"`
	Name      string    `validate:"required"`
	Email     string    `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email string) (*User, error) {
	validate := validator.New()

	user := User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := validate.Struct(user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
