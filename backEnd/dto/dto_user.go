package dto

type UserDto struct {
	UserId   int    `json:"id"`
	Name     string `json:"user_name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Number   int    `json:"number"`
}

type UsersDto []UserDto
