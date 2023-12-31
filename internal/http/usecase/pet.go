package usecase

import (
	"errors"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
)

type PetUseCase interface {
	CreatePet(pet model.Pet) (model.Pet, error)
	UpdatePet(pet model.Pet) (model.Pet, error)
	DeletePet(petID uint) (error)
}

func (uc *UseCase) CreatePet(pet model.Pet) (model.Pet, error) {
	if pet.Name == "" {
		return model.Pet{}, errors.New("название должно быть заполнено")
	}

	if pet.Birthdate == "" {
		return model.Pet{}, errors.New("дата рождения должна быть заполнена")
	}
	if pet.Spicies == "" {
		return model.Pet{}, errors.New("вид животного должен быть заполнен")
	}
	// Другие проверки на валидность полей животного добавим позже

	err := uc.Repository.CreatePet(pet)
	if err != nil {
		return model.Pet{}, errors.New("ошибка при создании питомца")
	}

	return pet, nil
}

func (uc *UseCase) UpdatePet(pet model.PetUpdateRequest) (model.PetUpdateRequest, error) {
	if pet.ID == 0 {
		return model.PetUpdateRequest{}, errors.New("неверный идентификатор питомца")
	}

	// Другие проверки на валидность полей животного добавим позже

	err := uc.Repository.UpdatePet(pet)
	if err != nil {
		return model.PetUpdateRequest{}, errors.New("ошибка при обновлении информации о питомце")
	}

	return pet, nil
}

func (uc *UseCase) DeletePet(petID uint) error {
	if petID == 0 {
		return errors.New("неверный идентификатор питомца")
	}

	// Другие проверки на валидность полей животного добавим позже
	
	err := uc.Repository.DeletePet(petID)
	if err != nil {
		return errors.New("ошибка при удалении питомца")
	}

	return nil
}

