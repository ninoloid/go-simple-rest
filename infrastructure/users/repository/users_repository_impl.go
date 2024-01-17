package user_repository

import (
	"context"
	"errors"

	users "github.com/ninoloid/go-simple-rest/domain/users/entities"
	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	DB *gorm.DB
}

func NewUsersRepository(DB *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{
		DB: DB,
	}
}

func (repository *UsersRepositoryImpl) FindAll(ctx context.Context) []users.User {
	var userList []users.User
	repository.DB.Find(&userList)

	return userList
}

func (repository *UsersRepositoryImpl) FindById(ctx context.Context, userId int) (users.User, error) {
	var user users.User
	if err := repository.DB.First(&user, userId).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return user, errors.New("User is not found")
		default:
			return user, errors.New("Internal error")
		}
	}

	return user, nil
}
