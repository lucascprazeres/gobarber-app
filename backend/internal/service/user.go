package service

import (
	"context"
	"gobarber/internal/domain/model"
	userRepo "gobarber/internal/infra/database/repository"
	"gobarber/internal/schema"
	"gobarber/pkg/errorwrapper"
	"net/http"
)

const (
	UnknownError           = "unknown-error"
	EmailAlreadyInUseError = "email-already-in-use"
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
		return nil, errorwrapper.New(UnknownError, err.Error())
	}

	if foundUser != nil {
		return nil, errorwrapper.New(EmailAlreadyInUseError, "Esse email já está sendo utilizado.").
			WithStatus(http.StatusConflict)
	}

	usr := model.FromSchemaInput(input)
	if err := s.userRepo.CreateUser(ctx, usr); err != nil {
		return nil, errorwrapper.New(UnknownError, err.Error())
	}

	return &schema.CreateUserOutput{ID: usr.ID.String()}, nil
}
