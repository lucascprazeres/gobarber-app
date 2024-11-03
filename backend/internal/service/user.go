package service

import (
	"context"
	"errors"
	"gobarber/internal/domain/model"
	userRepo "gobarber/internal/infra/database/repository"
	"gobarber/internal/schema"
)

type UserService interface {
	CreateUser(ctx context.Context, input *schema.CreateUserInput) (*schema.CreateUserOutput, error)
}

type userService struct {
	userRepo userRepo.UserRepository
}

func NewService() UserService {
	return &userService{
		userRepo: userRepo.NewRepository(),
	}
}

func (s *userService) CreateUser(ctx context.Context, input *schema.CreateUserInput) (*schema.CreateUserOutput, error) {
	foundUser, err := s.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if foundUser != nil {
		return nil, errors.New("email already in use")
	}

	usr := model.FromSchemaInput(input)

	if err := s.userRepo.CreateUser(ctx, usr); err != nil {
		return nil, err
	}

	return &schema.CreateUserOutput{ID: usr.ID.String()}, nil
}
