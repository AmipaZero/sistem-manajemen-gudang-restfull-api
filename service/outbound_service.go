package service

import (
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/repository"
	"sistem-manajemen-gudang/exception"
	"time"
)

type OutboundService interface {
	Create(p domain.Outbound) (domain.Outbound, error)
	GetAll() ([]domain.Outbound, error)
	GetByID(id uint) (domain.Outbound, error)
	Update(id uint, supplier string, receivedAt time.Time) (domain.Outbound, error)
	Delete(id uint) error

}

type outboundService struct {
	repo repository.OutboundRepository
}

func NewOutboundService(r repository.OutboundRepository) OutboundService {
	return &outboundService{repo: r}
}

func (s *outboundService) Create(i domain.Outbound) (domain.Outbound, error) {
    //  Ambil produk berdasarkan 
    product, err := s.repo.FindProductByID(i.ProductID)
    if err != nil {
        return i, err
    }

    //  Cek stok cukup
    if product.Stock < i.Quantity {
        return i, exception.ErrOutboundAdd
    }

    // Kurangi stok
    product.Stock -= i.Quantity

    // Update stok produk
    if err := s.repo.UpdateProduct(product); err != nil {
        return i, err
    }

    // Simpan record outbound
    outbound, err := s.repo.Create(i)
    if err != nil {
        return i, err
    }

    return outbound, nil
}

func (s *outboundService) GetAll() ([]domain.Outbound, error) {

	return s.repo.FindAll()
}
func (s *outboundService) GetByID(id uint) (domain.Outbound, error) {
	return s.repo.FindByID(id)
}

func (s *outboundService) Update(id uint, destination string, sent_at time.Time) (domain.Outbound, error) {

    // Ambil outbound 
    outbound, err := s.repo.FindByID(id)
    if err != nil {
        return outbound, err
    }

    // Hanya update field
    outbound.Destination = destination
    outbound.SentAt = sent_at
	

    //ke repository 
    updated, err := s.repo.Update(outbound)
    if err != nil {
        return updated, err
    }

    return updated, nil
}


func (s *outboundService) Delete(id uint) error {
	return s.repo.Delete(id)
}

