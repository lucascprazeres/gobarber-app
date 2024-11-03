package repository

import (
	"context"
	"errors"
	"gobarber/internal/domain/model"
	"gobarber/internal/infra/database"
	"gobarber/internal/infra/database/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewRepository() UserRepository {
	return &userRepository{
		DB: database.GetConnection(),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	userEntity := entity.FromUserModel(user)
	return r.DB.WithContext(ctx).Create(userEntity).Error
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var usr model.User
	if err := r.DB.WithContext(ctx).First(&usr, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &usr, nil
}
