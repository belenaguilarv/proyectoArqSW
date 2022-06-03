package dto

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"user_name"`
	Password string `json:"password"`
}

type UsersDto []UserDto
