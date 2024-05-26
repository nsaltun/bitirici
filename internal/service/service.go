package service

import "github.com/nsaltun/bitirici/internal/model"

type UserService interface {
	CreateUser(user model.User) (model.User, error)
	GetUsers() ([]model.User, error)
	GetUser(id string) (model.User, error)
	UpdateUser(id string, updateUser model.User) (model.User, error)
	DeleteUser(id string) error
}
