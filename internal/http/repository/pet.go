package repository

import (
	"errors"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
)

type PetRepository interface {
	CreatePet(userID uint64, requestPet model.Pet) (model.Pet, error)
	GetPetById(petID uint64) error
	GetAllPets() ([]model.Pet, error)
	DeletePet(petID uint64) error
	UpdatePet(petID uint64, pet model.Pet) (model.Pet, error)
}

func (r *Repository) CreatePet(userID uint64, pet model.Pet) (model.Pet, error) {
	if err := r.db.Table("pet").Create(&pet).Error; err != nil {
		return model.Pet{}, errors.New("ошибка при создании питомца и добавлении его в БД")
	}

	userPet := model.UserPet{
		UserID: userID,
		PetID:  pet.ID,
	}

	if err := r.db.Table("users_pets").Create(&userPet).Error; err != nil {
		r.db.Table("pet").Delete(&pet)
		return model.Pet{}, errors.New("ошибка при создании связи между пользователем и питомцем")
	}

	return pet, nil
}

func (r *Repository) GetAllPets(userID uint64) ([]model.Pet, error) {
    var pets []model.Pet
    if err := r.db.Table("pet").
		Joins("JOIN users_pets ON pet.id = users_pets.pet_id").
		Where("users_pets.user_id = ? AND pet.status = ?", userID, model.PET_STATUS_ACTIVE).
		Find(&pets).Error; err != nil {
		return nil, errors.New("ошибка при получении списка питомцев пользователя из БД")
	}
	return pets, nil
}

func (r *Repository) GetPetById(userID, petID uint64) (model.Pet, error) {
	var pet model.Pet
	if err := r.db.Table("pet").Where("id = ? AND status = ?", petID, model.PET_STATUS_ACTIVE).First(&pet).Error; err != nil {
		return model.Pet{}, errors.New("ошибка при получении активного питомца из БД")
	}

	var userPet model.UserPet
	if err := r.db.Table("users_pets").Where("user_id = ? AND pet_id = ?", userID, petID).First(&userPet).Error; err != nil {
		return model.Pet{}, errors.New("питомец не найден у данного пользователя")
	}

	return pet, nil
}

func (r *Repository) UpdatePet(userID, petID uint64, pet model.PetUpdateRequest) (model.Pet, error) {
	var userPet model.UserPet
	if err := r.db.Table("users_pets").Where("user_id = ? AND pet_id = ?", userID, petID).First(&userPet).Error; err != nil {
		return model.Pet{}, errors.New("питомец не найден или не принадлежит данному пользователю")
	}

	if err := r.db.Table("pet").Model(&model.Pet{}).Where("id = ? AND status = ?", petID, model.PET_STATUS_ACTIVE).Updates(pet).Error; err != nil {
		return model.Pet{}, errors.New("ошибка при обновлении информации о питомце в БД")
	}

	var updatedPet model.Pet
	if err := r.db.Table("pet").Where("id = ? AND status = ?", petID, model.PET_STATUS_ACTIVE).First(&updatedPet).Error; err != nil {
		return model.Pet{}, errors.New("ошибка при получении обновленной информации о питомце из БД")
	}

	return updatedPet, nil
}


func (r *Repository) DeletePet(userID, petID uint64) error {
	pet := &model.Pet{}
	if err := r.db.Table("pet").Where("id = ? AND status = ?", petID, model.PET_STATUS_ACTIVE).First(pet).Error; err != nil {
		return errors.New("питомец не найден или уже удален")
	}

	pet.Status = model.PET_STATUS_DELETED
	if err := r.db.Table("pet").Save(pet).Error; err != nil {
		return errors.New("ошибка при обновлении статуса питомца в БД")
	}

	return nil
}
