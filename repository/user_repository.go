package repository

import (
	"sistem-manajemen-gudang/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *model.User) error
	// FindByUsername(username string) (*model.User, error)
	// UpdateToken(id uint, token *string) error
	// FindById(id uint) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user *model.User) error {
	return r.db.Create(user).Error
}

// func (r *userRepository) FindByUsername(username string) (*model.User, error) {
// 	var user model.User
// 	err := r.db.Where("username = ?", username).First(&user).Error
// 	return &user, err
// }

// func (r *userRepository) UpdateToken(id uint, token *string) error {
// 	return r.db.Model(&model.User{}).Where("id = ?", id).Update("token", token).Error
// }

// func (r *userRepository) FindById(id uint) (*model.User, error) {
// 	var user model.User
// 	err := r.db.First(&user, id).Error
// 	return &user, err
// }
