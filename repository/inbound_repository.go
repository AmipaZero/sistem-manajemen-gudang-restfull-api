
package repository

import (
	"sistem-manajemen-gudang/model"
	"gorm.io/gorm"
)

type InboundRepository interface {
	Create(i model.Inbound) (model.Inbound, error)
	FindAll() ([]model.Inbound, error)
	FindByID(id uint) (model.Inbound, error)
	Update(i model.Inbound) (model.Inbound, error)
	Delete(id uint) error
	GetInboundLaporan(start, end string) ([]model.Inbound, error)
}

type inboundRepository struct {
	db *gorm.DB
}

func NewInboundRepository(db *gorm.DB) InboundRepository {
	return &inboundRepository{db: db}
}

func (r *inboundRepository) Create(i model.Inbound) (model.Inbound, error) {
	err := r.db.Create(&i).Error
	if err != nil {
		return i, err
	}
	// preload product setelah disimpan
	err = r.db.Preload("Product").First(&i, i.ID).Error
	return i, err
}

func (r *inboundRepository) FindAll() ([]model.Inbound, error) {
	var inbounds []model.Inbound
	err := r.db.Preload("Product").Find(&inbounds).Error
	return inbounds, err
}

func (r *inboundRepository) FindByID(id uint) (model.Inbound, error) {
	var inbound model.Inbound
	err := r.db.Preload("Product").First(&inbound, id).Error
	return inbound, err
}

func (r *inboundRepository) Update(i model.Inbound) (model.Inbound, error) {
	err := r.db.Model(&i).Updates(i).Error
	if err != nil {
		return i, err
	}
	// preload setelah update
	err = r.db.Preload("Product").First(&i, i.ID).Error
	return i, err
}

func (r *inboundRepository) Delete(id uint) error {
	return r.db.Delete(&model.Inbound{}, id).Error
}

func (r *inboundRepository) GetInboundLaporan(start, end string) ([]model.Inbound, error) {
	var inbounds []model.Inbound
	db := r.db.Preload("Product")

	if start != "" && end != "" {
		db = db.Where("tanggal BETWEEN ? AND ?", start, end)
	}

	err := db.Find(&inbounds).Error
	return inbounds, err
}
