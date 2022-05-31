package services

import (
	userCliente "go/clients" // conecta a la base de datos
	"go/dto"
	e "go/errors" // que hace esto? (utils)
	"go/model"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.UserName = user.UserName
	userDto.UserId = user.Id
	userDto.Password = user.Password
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers() // userClient??
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.UserName = user.Name
		userDto.Id = user.Id
		userDto.Password = user.Password

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}
