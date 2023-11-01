package model

import "time"

type Pet struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Species     string    `json:"species"`
	Breed       string    `json:"breed"`
	Birthdate   string    `json:"birthdate"`
	Gender      string    `json:"gender"`
	Color       string    `json:"color"`
	Weight      float64   `json:"weight"`
	Description string    `json:"description"`
	OwnerID     uint      `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	Photos      []string  `json:"photos"`
}
