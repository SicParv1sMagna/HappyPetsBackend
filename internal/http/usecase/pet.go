package usecase

import (
	"errors"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
)

type PetUseCase interface {
	CreatePet(userID uint64, requestPet model.PetCreateRequest) (model.Pet, error)
	GetAllPets(userID uint64) ([]model.Pet, error)
	GetPetById(userID, petID uint64) (model.Pet, error)
	UpdatePet(userID, petID uint64, pet model.PetUpdateRequest) (model.Pet, error)
	DeletePet(userID, petID uint64) error
}

func (uc *UseCase) CreatePet(userID uint64, requestPet model.PetCreateRequest) (model.Pet, error) {
	if requestPet.Name == "" {
		return model.Pet{}, errors.New("название должно быть заполнено")
	}
	if requestPet.Birthday.IsZero() {
		return model.Pet{}, errors.New("дата рождения должна быть заполнена")
	}
	if requestPet.Spicies == "" {
		return model.Pet{}, errors.New("вид животного должен быть заполнен")
	}
	if requestPet.Gender == "" {
		return model.Pet{}, errors.New("пол животного должен быть заполнен")
	}
	// Другие проверки на валидность полей животного добавим позже

	pet := model.Pet{
		Name:     requestPet.Name,
		Birthday: requestPet.Birthday,
		Spicies:  requestPet.Spicies,
		Age:      requestPet.Age,
		Gender:   requestPet.Gender,
		Color:    requestPet.Color,
		Weight:   requestPet.Weight,
		Food:     requestPet.Food,
		LivesAt:  requestPet.LivesAt,
		Status:   model.PET_STATUS_ACTIVE,
	}

	createdPet, err := uc.Repository.CreatePet(userID, pet)
	if err != nil {
		return model.Pet{}, err
	}

	return createdPet, nil
}

func (uc *UseCase) GetAllPets(userID uint64) ([]model.Pet, error) {
	pets, err := uc.Repository.GetAllPets(userID)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (uc *UseCase) GetPetById(userID, petID uint64) (model.Pet, error) {
	pet, err := uc.Repository.GetPetById(userID, petID)
	if err != nil {
		return model.Pet{}, err
	}
	return pet, nil
}

func (uc *UseCase) UpdatePet(userID, petID uint64, pet model.PetUpdateRequest) (model.Pet, error) {	
	if petID == 0 {
		return model.Pet{}, errors.New("неверный идентификатор питомца")
	}
	// Другие проверки на валидность полей животного добавим позже

	updatedPet, err := uc.Repository.UpdatePet(userID, petID, pet)
	if err != nil {
		return model.Pet{}, err
	}

	return updatedPet, nil
}

func (uc *UseCase) DeletePet(userID, petID uint64) error {
	if petID == 0 {
		return errors.New("неверный идентификатор питомца")
	}
	
	err := uc.Repository.RemoveServiceImage(userID, petID)
	if err != nil {
		return err
	}

	err = uc.Repository.DeletePet(userID, petID)
	if err != nil {
		return err
	}

	return nil
}

