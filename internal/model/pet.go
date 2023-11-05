package model

import (
	"time"
)

type Pet struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Birthdate   string    `json:"birthdate"`
	Gender      string    `json:"gender"`
	Color       string    `json:"color"`
	Weight      float64   `json:"weight"`
	Description string    `json:"description"`
	Spicies     string    `json:"spicies"`
	Status      string    `json:"status"`
	Food        string    `json:"food"`
	LivesAt     string    `json:"lives_at"`
	BreedID     uint64    `json:"breed_id"`
	TodoID      uint64    `json:"todo_id"`
	PassportID  uint64    `json:"passport_id"`
	CreatedAt   time.Time `json:"created_at"`
	Photos      []string  `json:"photos" gorm:"-"`
}

type PetUpdateRequest struct {
	ID          uint64   `json:"id"`
	Name        string   `json:"name"`
	Birthdate   string   `json:"birthdate"`
	Gender      string   `json:"gender"`
	Color       string   `json:"color"`
	Weight      float64  `json:"weight"`
	Description string   `json:"description"`
	Spicies     string   `json:"spicies"`
	Food        string   `json:"food"`
	Photos      []string `json:"photos"`
}
