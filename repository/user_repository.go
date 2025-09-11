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


