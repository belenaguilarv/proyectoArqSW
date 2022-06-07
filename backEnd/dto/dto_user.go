package dto

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"user_name"`
	Password string `json:"password"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Number   int    `json:"number"`
}

type UsersDto []UserDto
