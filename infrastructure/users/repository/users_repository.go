package user_repository

import (
	"context"

	users "github.com/ninoloid/go-simple-rest/domain/users/entities"
)

type UsersRepository interface {
	FindAll(ctx context.Context) []users.User
	FindById(ctx context.Context, categoryId int) (users.User, error)
}
