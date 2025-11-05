package service

import (
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/repository"
)

type OutboundService interface {
	Create(p domain.Outbound) (domain.Outbound, error)
	GetAll() ([]domain.Outbound, error)
	GetByID(id uint) (domain.Outbound, error)
	Update(p domain.Outbound) (domain.Outbound, error)
	Delete(id uint) error

}

type outboundService struct {
	repo repository.OutboundRepository
}

func NewOutboundService(r repository.OutboundRepository) OutboundService {
	return &outboundService{repo: r}
}

func (s *outboundService) Create(p domain.Outbound) (domain.Outbound, error) {
	return s.repo.Create(p)
}

func (s *outboundService) GetAll() ([]domain.Outbound, error) {

	return s.repo.FindAll()
}
func (s *outboundService) GetByID(id uint) (domain.Outbound, error) {
	return s.repo.FindByID(id)
}

func (s *outboundService) Update(p domain.Outbound) (domain.Outbound, error) {

	return s.repo.Update(p)
}

func (s *outboundService) Delete(id uint) error {
	return s.repo.Delete(id)
}

