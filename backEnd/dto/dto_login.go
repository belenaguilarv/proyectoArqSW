package dto

type LoginDto struct {
	Name     string `json:"user_name"`
	Password string `json:"password"`
}

type LoginsDto []LoginDto
