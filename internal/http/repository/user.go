package repository

import "github.com/SicParv1sMagna/HappyPetsBackend/internal/model"

type UserRepository interface {
	Create(user model.User) error
	GetById(id uint) error
	Delete(id uint) error
	Update(user model.User) error
}

func (r *Repository) Create(user model.User) error {
	return nil
}
