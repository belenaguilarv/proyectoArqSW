package dto

type UserDto struct {
<<<<<<< HEAD
	//UserId   int    `json:"id"`
=======
>>>>>>> 973f972940c6aa743ddeacd2ce2a382ca1df3a56
	Id       int    `json:"id"`
	Name     string `json:"user_name"`
	Password string `json:"password"`
	//Token    string `json:"token"` // para encryptar password
	//City     string `json:"city"`
	//	Street   string `json:"street"`
	//	Number   int    `json:"number"`
}

type UsersDto []UserDto
