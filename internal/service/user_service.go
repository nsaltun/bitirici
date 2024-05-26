// services/user_service.go
package service

import (
	"github.com/nsaltun/bitirici/internal/model"
	"github.com/nsaltun/bitirici/internal/repository"
)

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{UserRepository: userRepository}
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
	return s.UserRepository.Create(user)
}

func (s *userService) GetUsers() ([]model.User, error) {
	return s.UserRepository.GetAll()
}

func (s *userService) GetUser(id string) (model.User, error) {
	return s.UserRepository.GetByID(id)
}

func (s *userService) UpdateUser(id string, updateUser model.User) (model.User, error) {
	return s.UserRepository.Update(id, updateUser)
}

func (s *userService) DeleteUser(id string) error {
	return s.UserRepository.Delete(id)
}
