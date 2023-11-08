package model

type Passport struct {
	ID         uint64 `json:"id" gorm:"primarykey;autoIncrement"`
	HasChip    bool   `json:"has_chip"`
	Sterilized bool   `json:"sterilized"`
	PetID      uint   `json:"pet_id" gorm:"column:pet_id"`
}
