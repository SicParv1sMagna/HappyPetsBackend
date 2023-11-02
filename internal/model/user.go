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

type LoginUserRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
