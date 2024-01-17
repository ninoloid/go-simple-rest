package repository

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
	// SQL := "select id, name from users"
	// rows, err := tx.QueryContext(ctx, SQL)
	// helper.PanicIfError(err)
	// defer rows.Close()

	// var userList []users.User
	// for rows.Next() {
	// 	user := users.User{}
	// 	err := rows.Scan(&user.Id, &user.FirstName, &user.LastName)
	// 	helper.PanicIfError(err)
	// 	userList = append(userList, user)
	// }
	return userList
}

func (repository *UsersRepositoryImpl) FindById(ctx context.Context, userId int) (users.User, error) {
	// SQL := "select id, name from users where id = ?"
	// rows, err := tx.QueryContext(ctx, SQL, userId)
	// helper.PanicIfError(err)
	// defer rows.Close()

	// user := users.User{}
	// if rows.Next() {
	// 	err := rows.Scan(&user.Id, &user.FirstName, &user.LastName)
	// 	helper.PanicIfError(err)
	// 	return user, nil
	// } else {
	// 	return user, errors.New("user is not found")
	// }
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
