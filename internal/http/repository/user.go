package repository

import "github.com/SicParv1sMagna/HappyPetsBackend/internal/model"

type UserRepository interface {
	CreateUser(user model.User) error
	GetUserById(id uint) error
	DeleteUser(id uint) error
	UpdateUser(user model.User) error
	GetByPhoneNumber(phoneNumber string) (model.User, error)
}

func (r *Repository) CreateUser(user model.User) error {
	err := r.db.Table("User").Create(&user).Error
	return err
}

func (r *Repository) GetByPhoneNumber(phoneNumber string) (model.User, error) {
	var user model.User

	err := r.db.Table("User").Where(`"phone_number" = ?`, phoneNumber).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserById(id uint) (model.User, error) {
	var user model.User

	err := r.db.Table("User").Where(`"id" = ?`, id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
