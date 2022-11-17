package products

import (
	"3_2/internal/domain"
)

type Service interface {
	Store(domain.Product) (int, error)
	GetByName(name string) (domain.Product, error)
	GetAll(int) ([]domain.Product, error)
	Delete(id int) (error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Store(p domain.Product) (int, error) {
	return s.repo.Store(p)
}

func (s *service) GetByName(name string) (domain.Product, error) {
	return s.repo.GetByName(name)
}

func (s *service) GetAll(id int) ([]domain.Product, error) {
	return s.repo.GetAll(id)
}

func (s *service) Delete(id int) (error) {
	return s.repo.Delete(id)
}
