package repository

import (
	"gorm.io/gorm"
	"kolaborasi/entity"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	FindUserByEmail(email string) (entity.User, error)
	FindUserById(ID int) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}

func (r *userRepository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email=?", email).Find(&user)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}

func (r *userRepository) FindUserById(ID int) (entity.User, error) {
	var user entity.User
	err := r.db.Find(&user, ID)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	err := r.db.Save(&user)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}
