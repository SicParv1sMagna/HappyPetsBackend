package repository

import (
	"strconv"
	"time"

	"github.com/SicParv1sMagna/HappyPetsBackend/internal/model"
)

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

	err := r.db.Table("User").Where(`"phone_number" = ?`, phoneNumber).Find(&user).Error
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

func (r *Repository) UpdateUser(user model.User) error {
	err := r.db.Table("User").Where(`"id" = ?`, user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SaveJWTToken(id uint, token string) error {
	expiration := 7 * 24 * time.Hour

	idStr := strconv.FormatUint(uint64(id), 10)

	err := r.rd.Set(idStr, token, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}
