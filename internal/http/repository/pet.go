package repository

import (
	"errors"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
)

type PetRepository interface {
	CreatePet(pet model.Pet) error
	GetPetById(id uint) error
	DeletePet(id uint) error
	UpdatePet(pet model.Pet) error
}

func (r*Repository) CreatePet(pet model.Pet) error{
	if err := r.db.Create(&pet).Error; err != nil {
		return errors.New("ошибка при создании питомца и добавлении его в БД")
	}
	return nil
}