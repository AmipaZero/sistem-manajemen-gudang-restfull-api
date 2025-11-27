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
	Delete(id uint) error

	FindProductByID(id uint) (domain.Product, error)
	UpdateProduct(p domain.Product) error
}

type outboundRepository struct {
	db *gorm.DB
}

func NewOutboundRepository(db *gorm.DB) OutboundRepository {
	return &outboundRepository{db: db}
}

func (r *outboundRepository) Create(i domain.Outbound) (domain.Outbound, error) {
	if err := r.db.Create(&i).Error; err != nil {
		return i, err
	}

	// preload product 
	err := r.db.Preload("Product").First(&i, i.ID).Error
	return i, err
}

func (r *outboundRepository) FindAll() ([]domain.Outbound, error) {
	var results []domain.Outbound
	err := r.db.Preload("Product").Find(&results).Error
	return results, err
}

func (r *outboundRepository) FindByID(id uint) (domain.Outbound, error) {
	var outbound domain.Outbound
	err := r.db.Preload("Product").First(&outbound, id).Error
	return outbound, err
}

func (r *outboundRepository) Update(i domain.Outbound) (domain.Outbound, error) {
    err := r.db.Model(&domain.Outbound{}).
        Where("id = ?", i.ID).
        Updates(map[string]interface{}{
            "sent_at":    i.SentAt,
            "destination": i.Destination,
        }).Error

    if err != nil {
        return i, err
    }

    // Ambil ulang data 
    r.db.Preload("Product").First(&i, i.ID)
    return i, nil
}

func (r *outboundRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Outbound{}, id).Error
}

func (r *outboundRepository) FindProductByID(id uint) (domain.Product, error) {
	var p domain.Product
	err := r.db.First(&p, id).Error
	return p, err
}

func (r *outboundRepository) UpdateProduct(p domain.Product) error {
	return r.db.Model(&p).Updates(p).Error
}
