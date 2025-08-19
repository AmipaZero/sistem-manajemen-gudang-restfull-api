package service

import (
	"sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/repository"
)

type InboundService interface {
	Create(p model.Inbound) (model.Inbound, error)
	GetAll() ([]model.Inbound, error)
	GetByID(id uint) (model.Inbound, error)
	Update(p model.Inbound) (model.Inbound, error)
	Delete(id uint) error
	GetLaporan(start, end string) ([]model.Inbound, error)
}

type inboundService struct {
	repo repository.InboundRepository
}

func NewInboundService(r repository.InboundRepository) InboundService {
	return &inboundService{repo: r}
}

func (s *inboundService) Create(p model.Inbound) (model.Inbound, error) {
	return s.repo.Create(p)
}

func (s *inboundService) GetAll() ([]model.Inbound, error) {

	return s.repo.FindAll()
}
func (s *inboundService) GetByID(id uint) (model.Inbound, error) {
	return s.repo.FindByID(id)
}

func (s *inboundService) Update(p model.Inbound) (model.Inbound, error) {

	return s.repo.Update(p)
}

func (s *inboundService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *inboundService) GetLaporan(start, end string) ([]model.Inbound, error) {
	return s.repo.GetInboundLaporan(start, end)
}
