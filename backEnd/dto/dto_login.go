package dto

type LoginDto struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type LoginsDto []LoginDto
