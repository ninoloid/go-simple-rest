package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/ninoloid/go-simple-rest/application/users/response"
	"github.com/ninoloid/go-simple-rest/common/helper"
	"github.com/ninoloid/go-simple-rest/infrastructure/users/repository"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersService(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

func (service *UsersServiceImpl) FindAll(ctx context.Context) []response.UsersResponse {

	userList := service.UsersRepository.FindAll(ctx)

	return helper.ToUsersResponses(userList)
}

func (service *UsersServiceImpl) FindById(ctx context.Context, userId int) (response.UsersResponse, error) {
	user, err := service.UsersRepository.FindById(ctx, userId)
	if err != nil {
		// panic(exception.NewNotFoundError(err.Error()))
		return helper.ToUsersResponse(user), err
	}

	return helper.ToUsersResponse(user), nil
}
