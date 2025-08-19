package service

import (
	"sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/repository"
)

type OutboundService interface {
	Create(p model.Outbound) (model.Outbound, error)
	GetAll() ([]model.Outbound, error)
	GetByID(id uint) (model.Outbound, error)
	Update(p model.Outbound) (model.Outbound, error)
	Delete(id uint) error
	GetLaporan(start, end string) ([]model.Outbound, error)

}

type outboundService struct {
	repo repository.OutboundRepository
}

func NewOutboundService(r repository.OutboundRepository) OutboundService {
	return &outboundService{repo: r}
}

func (s *outboundService) Create(p model.Outbound) (model.Outbound, error) {
	return s.repo.Create(p)
}

func (s *outboundService) GetAll() ([]model.Outbound, error) {

	return s.repo.FindAll()
}
func (s *outboundService) GetByID(id uint) (model.Outbound, error) {
	return s.repo.FindByID(id)
}

func (s *outboundService) Update(p model.Outbound) (model.Outbound, error) {

	return s.repo.Update(p)
}

func (s *outboundService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *outboundService) GetLaporan(start, end string) ([]model.Outbound, error) {
	return s.repo.GetOutboundLaporan(start, end)
}