package model

type Breed struct {
	ID          uint64 `json:"id" gorm:"primarykey;autoIncrement"`
	Species     string `json:"species"`
	Description string `json:"description"`
	Allowed     []byte `json:"allowed"`
	NotAllowed  []byte `json:"not_allowed"`
}