package dto

type UserDto struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"user_password"`
	}

type UsersDto []UserDto
