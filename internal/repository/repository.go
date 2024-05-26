package repository

import "github.com/nsaltun/bitirici/internal/model"

type UserRepository interface {
	Create(user model.User) (model.User, error)
	GetAll() ([]model.User, error)
	GetByID(id string) (model.User, error)
	Update(id string, updateUser model.User) (model.User, error)
	Delete(id string) error
}
