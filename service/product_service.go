package service

import (
	"sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/repository"
)

type ProductService interface {
	Create(p model.Product) (model.Product, error)
	GetAll() ([]model.Product, error)
	GetByID(id uint) (model.Product, error)
	Update(p model.Product) (model.Product, error)
	Delete(id uint) error
	GetLaporan(start, end string) ([]model.Product, error)

}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

func (s *productService) Create(p model.Product) (model.Product, error) {
	return s.repo.Save(p)
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.FindAll()
}
func (s *productService) GetByID(id uint) (model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Update(p model.Product) (model.Product, error) {
	return s.repo.Update(p)
}

func (s *productService) Delete(id uint) error {
	return s.repo.Delete(id)
}
func (s *productService) GetLaporan(start, end string) ([]model.Product, error) {
	return s.repo.GetProductLaporan(start, end)
}