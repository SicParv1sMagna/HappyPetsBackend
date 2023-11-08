package model

type UserPet struct {
	UserID uint64 `json:"user_id" gorm:"user_id"`
	PetID  uint64 `json:"pet_id" gorm:"pet_id"`
}