package model

type User struct {
	ID             uint64 `json:"id" gorm:"primarykey;autoIncrement"`
	FirstName      string `json:"first_name" gorm:"column:first_name"`
	LastName       string `json:"last_name" gorm:""column:second_name"`
	Email          string `json:"email" gorm:"column:email"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number"`
	Password       string `json:"password" gorm:"column:password"`
	RepeatPassword string `json:"repeat_password" gorm:"-"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserUpdateRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}
