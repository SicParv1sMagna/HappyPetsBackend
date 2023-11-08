package model

import "time"

type Todo struct {
	ID          uint64    `json:"id" gorm:"primarykey;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Closed      bool      `json:"closed"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
	PetID       uint      `json:"pet_id" gorm:"column:pet_id"`
}