package helper

import (
	"github.com/ninoloid/go-simple-rest/application/users/response"
	users "github.com/ninoloid/go-simple-rest/domain/users/entities"
)

func ToUsersResponse(user users.User) response.UsersResponse {
	return response.UsersResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToUsersResponses(userList []users.User) []response.UsersResponse {
	var userResponses []response.UsersResponse
	for _, user := range userList {
		userResponses = append(userResponses, ToUsersResponse(user))
	}
	return userResponses
}
