package repository

import (
	"sistem-manajemen-gudang/model/domain"

	"gorm.io/gorm"
)

type InboundRepository interface {
	Create(i domain.Inbound) (domain.Inbound, error)
	FindAll() ([]domain.Inbound, error)
	FindByID(id uint) (domain.Inbound, error)
	Update(i domain.Inbound) (domain.Inbound, error)
	Delete(id uint) error
}

type inboundRepository struct {
	db *gorm.DB
}

func NewInboundRepository(db *gorm.DB) InboundRepository {
	return &inboundRepository{db: db}
}

func (r *inboundRepository) Create(i domain.Inbound) (domain.Inbound, error) {
	if err := r.db.Create(&i).Error; err != nil {
		return i, err
	}
	r.db.Preload("Product").First(&i, i.ID)
	return i, nil
}

func (r *inboundRepository) FindAll() ([]domain.Inbound, error) {
	var result []domain.Inbound
	err := r.db.Preload("Product").Find(&result).Error
	return result, err
}

func (r *inboundRepository) FindByID(id uint) (domain.Inbound, error) {
	var inbound domain.Inbound
	err := r.db.Preload("Product").First(&inbound, id).Error
	return inbound, err
}

func (r *inboundRepository) Update(i domain.Inbound) (domain.Inbound, error) {
	err := r.db.Model(&i).Updates(i).Error
	if err != nil {
		return i, err
	}
	r.db.Preload("Product").First(&i, i.ID)
	return i, nil
}

func (r *inboundRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Inbound{}, id).Error
}

