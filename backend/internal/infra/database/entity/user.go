package entity

import (
	"github.com/google/uuid"
	"gobarber/internal/domain/model"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Avatar    *string
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null,unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

func FromUserModel(usr *model.User) *User {
	return &User{
		ID:       usr.ID,
		Avatar:   usr.Avatar,
		Name:     usr.Name,
		Email:    usr.Email,
		Password: usr.Password,
	}
}
