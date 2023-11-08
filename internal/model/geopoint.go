package model

import "time"

type Geopoint struct {
	ID          uint64    `json:"id" gorm:"primarykey;autoIncrement"`
	CreatedAt   time.Time `json:"created_at"`
	Phone       string    `json:"phone"`
	Description string    `json:"description"`
	Breed       string    `json:"breed"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	UserID      uint      `json:"user_id" gorm:"column:user_id"`
}