package repository

import (
	"sistem-manajemen-gudang/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(p model.Product) (model.Product, error)
	FindAll() ([]model.Product, error)
	FindByID(id uint) (model.Product, error)
	Update(p model.Product) (model.Product, error)
	Delete(id uint) error
	GetProductLaporan(start, end string) ([]model.Product, error)
	
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Save(product model.Product) (model.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepo) FindAll() ([]model.Product, error) {
	var people []model.Product
	err := r.db.Find(&people).Error
	return people, err
}
func (r *productRepo) FindByID(id uint) (model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *productRepo) Update(product model.Product) (model.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *productRepo) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

func (r *productRepo) GetProductLaporan(start, end string) ([]model.Product, error) {
	var products []model.Product
	db := r.db

	if start != "" && end != "" {
		db = db.Where("created_at BETWEEN ? AND ?", start, end)
	}

	err := db.Find(&products).Error
	return products, err
}
