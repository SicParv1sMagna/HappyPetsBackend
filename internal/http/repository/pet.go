package repository

import (
	"errors"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
)

type PetRepository interface {
	CreatePet(pet model.Pet) error
	GetPetById(petID uint) error
	DeletePet(petID uint) error
	UpdatePet(pet model.Pet) error
}

func (r *Repository) CreatePet(pet model.Pet) error {
	if err := r.db.Create(&pet).Error; err != nil {
		return errors.New("ошибка при создании питомца и добавлении его в БД")
	}
	return nil
}

func (r *Repository) GetPetById(petID uint) (model.Pet, error) {
	var pet model.Pet
	if err := r.db.First(&pet, petID).Error; err != nil {
		return model.Pet{}, errors.New("ошибка при получении питомца из БД")
	}
	return pet, nil
}

func (r *Repository) UpdatePet(pet model.PetUpdateRequest) error {
	if err := r.db.Model(&model.Pet{}).Where("id = ?", pet.ID).Updates(pet).Error; err != nil {
		return errors.New("ошибка при обновлении информации о питомце в БД")
	}
	return nil
}

func (r *Repository) DeletePet(petID uint) error {
	pet := &model.Pet{}
	if err := r.db.Where("id = ?", petID).First(pet).Error; err != nil {
		return errors.New("питомец не найден")
	}

	pet.Status = model.PET_STATUS_DELETED
	if err := r.db.Save(pet).Error; err != nil {
		return errors.New("ошибка при обновлении статуса питомца в БД")
	}

	return nil
}
