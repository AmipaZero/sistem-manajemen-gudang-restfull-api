package repository

import (
	"sistem-manajemen-gudang/model/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(p domain.Product) (domain.Product, error)
	FindAll() ([]domain.Product, error)
	FindByID(id uint) (domain.Product, error)
	Update(p domain.Product) (domain.Product, error)
	Delete(id uint) error

}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Save(product domain.Product) (domain.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepo) FindAll() ([]domain.Product, error) {
	var product []domain.Product
	err := r.db.Find(&product).Error
	return product, err
}
func (r *productRepo) FindByID(id uint) (domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *productRepo) Update(product domain.Product) (domain.Product, error) {
	err := r.db.Model(&domain.Product{}).
		Where("id = ?", product.ID).
		Updates(map[string]interface{}{
			"name":     product.Name,
			"sku":      product.SKU,
			"category": product.Category,
			"unit":     product.Unit,
		}).Error

	if err != nil {
		return product, err
	}

	r.db.First(&product, product.ID)
	return product, nil
}


func (r *productRepo) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}