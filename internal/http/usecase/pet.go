package usecase

import (
	"errors"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
)

type PetUseCase interface {
	CreatePet(pet model.Pet) (model.Pet, error)
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
