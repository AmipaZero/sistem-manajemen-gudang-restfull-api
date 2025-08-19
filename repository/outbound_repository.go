package repository

import (
	"sistem-manajemen-gudang/model"
	"gorm.io/gorm"
)

type OutboundRepository interface {
	Create(i model.Outbound) (model.Outbound, error)
	FindAll() ([]model.Outbound, error)
	FindByID(id uint) (model.Outbound, error)
	Update(i model.Outbound) (model.Outbound, error)
	Delete(i uint) error
	GetOutboundLaporan(start, end string) ([]model.Outbound, error)
}
type outboundRepository struct {
	db *gorm.DB
}

func NewOutboundRepository(db *gorm.DB) OutboundRepository {
	return &outboundRepository{db: db}
}

func (r *outboundRepository) Create(i model.Outbound)(model.Outbound, error){
	err := r.db.Create(&i).Error
	if err != nil{
		return i, err
	}
	// preload product setelah disimpan
	err = r.db.Preload("Product").First(&i, i.ID).Error
	return i, err
}
func (r *outboundRepository) FindAll() ([]model.Outbound, error){
	var outbound []model.Outbound
	err := r.db.Preload("Product").Find(&outbound).Error
	return outbound, err
}
func (r *outboundRepository) FindByID(id uint) (model.Outbound, error){
	var outbound model.Outbound
	err := r.db.Preload("Product").First(&outbound, id).Error
	return outbound, err
}
func (r *outboundRepository) Update(i model.Outbound) (model.Outbound, error){
	err := r.db.Model(&i).Updates(i).Error
	if err != nil{
		return i, err
	}
	err = r.db.Preload("Product").First(&i, i.ID).Error
	return i, err
}
func (r *outboundRepository) Delete(id uint) error{
	return r.db.Delete(&model.Outbound{}, id).Error
}
func (r *outboundRepository) GetOutboundLaporan(start, end string) ([]model.Outbound, error) {
	var outbound []model.Outbound
	db := r.db.Preload("Product")

	if start != "" && end != "" {
		db = db.Where("tanggal BETWEEN ? AND ?", start, end)
	}

	err := db.Find(&outbound).Error
	return outbound, err
}
