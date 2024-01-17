package user_service

import (
	"context"

	"github.com/ninoloid/go-simple-rest/application/users/response"
)

type UsersService interface {
	FindAll(ctx context.Context) []response.UsersResponse
	FindById(ctx context.Context, categoryId int) response.UsersResponse
}
