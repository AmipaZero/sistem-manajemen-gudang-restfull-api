package service

import (
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/repository"
)

type InboundService interface {
	Create(i domain.Inbound) (domain.Inbound, error)
	GetAll() ([]domain.Inbound, error)
	GetByID(id uint) (domain.Inbound, error)
	Update(i domain.Inbound) (domain.Inbound, error)
	Delete(id uint) error
}

type inboundService struct {
	repo repository.InboundRepository
}

func NewInboundService(r repository.InboundRepository) InboundService {
	return &inboundService{repo: r}
}

func (s *inboundService) Create(i domain.Inbound) (domain.Inbound, error) {
	return s.repo.Create(i)
}

func (s *inboundService) GetAll() ([]domain.Inbound, error) {
	return s.repo.FindAll()
}

func (s *inboundService) GetByID(id uint) (domain.Inbound, error) {
	return s.repo.FindByID(id)
}

func (s *inboundService) Update(i domain.Inbound) (domain.Inbound, error) {
	return s.repo.Update(i)
}

func (s *inboundService) Delete(id uint) error {
	return s.repo.Delete(id)
}


