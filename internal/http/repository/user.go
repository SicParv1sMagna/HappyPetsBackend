package repository

import "github.com/SicParv1sMagna/HappyPetsBackend/internal/model"

type UserRepository interface {
	CreateUser(user model.User) error
	GetUserById(id uint) error
	DeleteUser(id uint) error
	UpdateUser(user model.User) error
}

func (r *Repository) CreateUser(user model.User) error {
	return nil
}
