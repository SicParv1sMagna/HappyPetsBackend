package repository

import (
	"fmt"

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
		return fmt.Errorf("ошибка при создании питомца и добавлении его в БД: %v", err)
	}
	return nil
}

func (r *Repository) GetPetById(petID uint) (model.Pet, error) {
	var pet model.Pet
	if err := r.db.First(&pet, petID).Error; err != nil {
		return model.Pet{}, fmt.Errorf("ошибка при получении питомца из БД: %v", err)
	}
	return pet, nil
}


func (r *Repository) UpdatePet(pet model.Pet) error {
	if err := r.db.Model(&model.Pet{}).Where("id = ?", pet.ID).Updates(pet).Error; err != nil {
		return fmt.Errorf("ошибка при обновлении информации о питомце в БД: %v", err)
	}
	return nil
}
