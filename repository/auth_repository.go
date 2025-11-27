package repository

import (
	"sistem-manajemen-gudang/model/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsername(username string) (*domain.User, error)
	UpdateToken(id uint, token *string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}


func (r *authRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *authRepository) UpdateToken(id uint, token *string) error {
	return r.db.Model(&domain.User{}).Where("id = ?", id).Update("token", token).Error
}

func (r *authRepository) FindById(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}
