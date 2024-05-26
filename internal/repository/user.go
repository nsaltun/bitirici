// repository/user_repository.go
package repository

import (
	"github.com/nsaltun/bitirici/internal/model"
	"github.com/nsaltun/bitirici/lib/db"
)

type userRepository struct {
	DB *db.PgDB
}

func NewUserRepository(db *db.PgDB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *userRepository) GetByID(id string) (model.User, error) {
	var user model.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) Update(id string, updateUser model.User) (model.User, error) {
	var user model.User
	err := r.DB.Model(&user).Where("id = ?", id).Updates(updateUser).Error
	return user, err
}

func (r *userRepository) Delete(id string) error {
	var user model.User
	err := r.DB.Where("id = ?", id).Delete(&user).Error
	return err
}
