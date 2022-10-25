package products

import (
	"clase3_parte2/internal/domain"
	"fmt"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(name, productType string, count int, price float64) (domain.Product, error)
	Delete(id int) (error)
	Update(id int, name string, typeProduct string, count int, price float64) (domain.Product, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]domain.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Store(name, productType string, count int, price float64) (domain.Product, error) {
	lastID, err := s.repo.LastID()
	if err != nil {
		return domain.Product{}, fmt.Errorf("error getting product last id: %w", err)
	}
	lastID++
	product, err := s.repo.Store(lastID, name, productType, count, price)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error creating product: %w", err)
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting product in %w", err)
	}
	return nil
}

func (s *service) Update(id int, name string, typeProduct string, count int, price float64) (domain.Product, error) {
	updatedProd, err := s.repo.Update(id, name, typeProduct, count, price)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product in: %w", err)
	}
	
	return updatedProd, nil
}
