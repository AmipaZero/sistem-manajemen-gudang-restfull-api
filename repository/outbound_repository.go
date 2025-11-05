package repository

import (
	"sistem-manajemen-gudang/model/domain"
	"gorm.io/gorm"
)

type OutboundRepository interface {
	Create(i domain.Outbound) (domain.Outbound, error)
	FindAll() ([]domain.Outbound, error)
	FindByID(id uint) (domain.Outbound, error)
	Update(i domain.Outbound) (domain.Outbound, error)
	Delete(i uint) error
}
type outboundRepository struct {
	db *gorm.DB
}

func NewOutboundRepository(db *gorm.DB) OutboundRepository {
	return &outboundRepository{db: db}
}

func (r *outboundRepository) Create(i domain.Outbound)(domain.Outbound, error){
	err := r.db.Create(&i).Error
	if err != nil{
		return i, err
	}
	// preload product setelah disimpan
	err = r.db.Preload("Product").First(&i, i.ID).Error
	return i, err
}
func (r *outboundRepository) FindAll() ([]domain.Outbound, error){
	var outbound []domain.Outbound
	err := r.db.Preload("Product").Find(&outbound).Error
	return outbound, err
}
func (r *outboundRepository) FindByID(id uint) (domain.Outbound, error){
	var outbound domain.Outbound
	err := r.db.Preload("Product").First(&outbound, id).Error
	return outbound, err
}
func (r *outboundRepository) Update(i domain.Outbound) (domain.Outbound, error){
	err := r.db.Model(&i).Updates(i).Error
	if err != nil{
		return i, err
	}
	err = r.db.Preload("Product").First(&i, i.ID).Error
	return i, err
}
func (r *outboundRepository) Delete(id uint) error{
	return r.db.Delete(&domain.Outbound{}, id).Error
}
