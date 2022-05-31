package dto

type UserDto struct {
	Id       int    `json:"user_id"`
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
}

type UsersDto []UserDto
