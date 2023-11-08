package model

import (
	"time"
)

const (
	PET_STATUS_ACTIVE  = "активный"
	PET_STATUS_DELETED = "удален"
)

type Pet struct {
	ID           uint64    `json:"id" gorm:"primarykey;autoIncrement"`
	Name         string    `json:"name"`
	Birthday     time.Time `json:"birthday"`
	Spicies      string    `json:"spicies"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	Color        string    `json:"color"`
	Weight       float64   `json:"weight"`
	CreationDate time.Time `json:"creation_date"`
	Lost         bool      `json:"lost"`
	Food         string    `json:"food"`
	LivesAt      string    `json:"lives_at"`
	BreedID      uint      `json:"breed_id" gorm:"column:breed_id"`
	Status       string    `json:"status"`
	Photo   	 string	   `json:"photo"`
}

type PetCreateRequest struct {
	Name         string    `json:"name"`
	Birthday     time.Time `json:"birthday"`
	Spicies      string    `json:"spicies"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	Color        string    `json:"color"`
	Weight       float64   `json:"weight"`
	Food         string    `json:"food"`
	LivesAt      string    `json:"lives_at"`
}

type PetUpdateRequest struct {
	Name         string    `json:"name"`
	Birthday     time.Time `json:"birthday"`
	Spicies      string    `json:"spicies"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	Color        string    `json:"color"`
	Weight       float64   `json:"weight"`
	Food         string    `json:"food"`
	LivesAt      string    `json:"lives_at"`
	Lost         bool      `json:"lost"`
}
