package usecase

import (
	"errors"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
	"github.com/SicParv1sMagna/HappyPetsBackend/internal/pkg/middleware"
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

	if userJSON.PhoneNumber == "" {
		return model.User{}, errors.New("номер телефона должен быть заполнен")
	}

	if userJSON.Email == "" {
		return model.User{}, errors.New("почта должна быть заполнена")
	}

	if userJSON.Password == "" {
		return model.User{}, errors.New("пароль должен быть заполнен")
	}

	if len(userJSON.Password) < 8 || len(userJSON.Password) > 20 {
		return model.User{}, errors.New("пароль должен содержать от 8 до 20 символов")
	}

	candidate, err := uc.Repository.GetByPhoneNumber(userJSON.PhoneNumber)
	if err != nil {
		return model.User{}, err
	}

	if candidate == userJSON {
		return model.User{}, errors.New("такой пользователь уже существует")
	}

	userJSON.Password, err = middleware.HashPassword(userJSON.Password)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		FirstName:   userJSON.FirstName,
		LastName:    userJSON.LastName,
		Email:       userJSON.Email,
		PhoneNumber: userJSON.PhoneNumber,
		Password:    userJSON.Password,
	}

	err := uc.Repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (uc *UseCase) LoginUser(userJSON model.LoginUserRequest) (string, error) {
	if userJSON.PhoneNumber == "" {
		return "", errors.New("запполните номер телефона")
	}

	if userJSON.Password == "" {
		return "", errors.New("заполните пароль")
	}

	candidate, err := uc.Repository.GetByPhoneNumber(userJSON.PhoneNumber)
	if err != nil {
		return "", err
	}

	if ok := middleware.CheckPasswordHash(userJSON.Password, candidate.Password); !ok {
		return "", errors.New("пароли не совпадают")
	}

	token, err := middleware.GenerateJWTToken(uint(candidate.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
