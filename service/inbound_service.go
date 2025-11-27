package service

import (
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/repository"
	"time"
)

type InboundService interface {
	Create(i domain.Inbound) (domain.Inbound, error)
	GetAll() ([]domain.Inbound, error)
	GetByID(id uint) (domain.Inbound, error)
	UpdateData(id uint, supplier string, receivedAt time.Time) (domain.Inbound, error)
	Delete(id uint) error
}

type inboundService struct {
	repo repository.InboundRepository
}

func NewInboundService(r repository.InboundRepository) InboundService {
	return &inboundService{repo: r}
}


func (s *inboundService) Create(i domain.Inbound) (domain.Inbound, error) {
    //  Ambil produk
    product, err := s.repo.FindProductByID(i.ProductID)
    if err != nil {
        return i, err
    }

    // Update stok produk
    product.Stock += i.Quantity

    // Simpan perubahan
    if err := s.repo.UpdateProduct(product); err != nil {
        return i, err
    }

    // Simpan 
    inbound, err := s.repo.Create(i)
    if err != nil {
        return i, err
    }

    return inbound, nil
}

func (s *inboundService) GetAll() ([]domain.Inbound, error) {
	return s.repo.FindAll()
}

func (s *inboundService) GetByID(id uint) (domain.Inbound, error) {
	return s.repo.FindByID(id)
}

func (s *inboundService) UpdateData(id uint, supplier string, receivedAt time.Time) (domain.Inbound, error) {

    //  Ambil inbound lama
    inbound, err := s.repo.FindByID(id)
    if err != nil {
        return inbound, err
    }

    // Hanya update field 
    inbound.Supplier = supplier
    inbound.ReceivedAt = receivedAt

    //  Kirim ke repository 
    updated, err := s.repo.Update(inbound)
    if err != nil {
        return updated, err
    }

    return updated, nil
}


func (s *inboundService) Delete(id uint) error {
	return s.repo.Delete(id)
}


