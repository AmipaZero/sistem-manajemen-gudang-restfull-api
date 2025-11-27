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
	FindProductByID(id uint) (domain.Product, error)
	UpdateProduct(p domain.Product) error

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
    err := r.db.Model(&domain.Inbound{}).
        Where("id = ?", i.ID).
        Updates(map[string]interface{}{
            "supplier":    i.Supplier,
            "received_at": i.ReceivedAt,
        }).Error

    if err != nil {
        return i, err
    }

    // Ambil ulang data setelah update
    r.db.Preload("Product").First(&i, i.ID)
    return i, nil
}


func (r *inboundRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Inbound{}, id).Error
}
func (r *inboundRepository) FindProductByID(id uint) (domain.Product, error) {
    var product domain.Product
    err := r.db.First(&product, id).Error
    return product, err
}

func (r *inboundRepository) UpdateProduct(p domain.Product) error {
    return r.db.Model(&p).Updates(p).Error
}

