package repository

import (
	"sistem-manajemen-gudang/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsername(username string) (*model.User, error)
	UpdateToken(id uint, token *string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

// func (r *userRepository) Save(user *model.User) error {
// 	return r.db.Create(user).Error
// }

func (r *authRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *authRepository) UpdateToken(id uint, token *string) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("token", token).Error
}

func (r *authRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}
