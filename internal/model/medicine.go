package model

import "time"

type Medicine struct {
	ID          uint64    `json:"id" gorm:"primarykey;autoIncrement"`
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
