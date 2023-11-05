package model

type User struct {
	ID             uint64 `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserUpdateRequest struct {
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	Password       string
	RepeatPassword string
}
