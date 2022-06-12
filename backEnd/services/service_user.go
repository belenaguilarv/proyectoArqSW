package services

import (
	userCliente "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/user"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"
	e "github.com/belenaguilarv/proyectoArqSW/backEnd/errors"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
	"github.com/golang-jwt/jwt"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	LoginUser(loginDto dto.LoginDto) (dto.Token, e.ApiError)
}

var jwtKey = []byte("secret_key")

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
	userDto.Name = user.Name
	userDto.Id = user.Id
	userDto.Password = user.Password
	userDto.Number = user.Number
	userDto.Street = user.Street
	userDto.City = user.City
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.Id = user.Id
		userDto.Password = user.Password
		userDto.Number = user.Number
		userDto.Street = user.Street
		userDto.City = user.City
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.Token, e.ApiError) {
	var user model.User = userCliente.GetUserByName(loginDto.Name)

	var tokenDto dto.Token

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("User not found")
	}

	if user.Password == loginDto.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.Id = user.Id
	}
	return tokenDto, nil
}
