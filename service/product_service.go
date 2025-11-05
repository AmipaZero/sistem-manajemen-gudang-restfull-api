package service

import (
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/repository"
)

type ProductService interface {
	Create(p domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetByID(id uint) (domain.Product, error)
	Update(p domain.Product) (domain.Product, error)
	Delete(id uint) error

}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

func (s *productService) Create(p domain.Product) (domain.Product, error) {
	return s.repo.Save(p)
}

func (s *productService) GetAll() ([]domain.Product, error) {
	return s.repo.FindAll()
}
func (s *productService) GetByID(id uint) (domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Update(p domain.Product) (domain.Product, error) {
	return s.repo.Update(p)
}

func (s *productService) Delete(id uint) error {
	return s.repo.Delete(id)
}
