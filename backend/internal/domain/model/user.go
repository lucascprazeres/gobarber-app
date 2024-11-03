package model

import (
	"github.com/google/uuid"
	"gobarber/internal/schema"
	"time"
)

type User struct {
	ID        uuid.UUID
	Avatar    *string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromSchemaInput(input *schema.CreateUserInput) *User {
	return &User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}
