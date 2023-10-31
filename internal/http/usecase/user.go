package usecase

import (
	"errors"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
)

type UserUseCase interface {
	RegisterUser(userJSON model.User) (model.User, error)
	LoginUser(userJSON model.User) (model.User, error)
}

func (uc *UseCase) RegisterUser(userJSON model.User) (model.User, error) {
	if userJSON.FirstName == "" {
		return model.User{}, errors.New("имя должно быть заполнено")
	}

	if userJSON.LastName == "" {
		return model.User{}, errors.New("фамилия должна быть заполнена")
	}

	if len(userJSON.Password) < 8 || len(userJSON.Password) > 20 {
		return model.User{}, errors.New("пароль должен содержать от 8 до 20 символов")
	}

	user := model.User{
		FirstName: userJSON.FirstName,
		LastName:  userJSON.LastName,
		Email:     userJSON.Email,
		Password:  userJSON.Password,
	}

	err := uc.Repository.Create(user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
