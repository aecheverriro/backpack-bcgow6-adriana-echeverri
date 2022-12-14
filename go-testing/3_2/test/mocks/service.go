package mocks

import (
	"clase3_parte2/internal/domain"
	"fmt"
)

type MockService struct {
	DataMock []domain.Product
	Error string
}

func (m *MockService) GetAll() ([]domain.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil

}

func (m *MockService) Store(name, productType string, count int, price float64) (domain.Product, error) {
	if m.Error != "" {
		return domain.Product{}, fmt.Errorf(m.Error)
	}
	p := domain.Product{
		Name: name,
		Type: productType,
		Count: count,
		Price: price,
	}
	return p, nil
}

func (m *MockService) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	if m.Error != "" {
		return domain.Product{}, fmt.Errorf(m.Error)
	}
	p := domain.Product{
		Name: name,
		Type: productType,
		Count: count,
		Price: price,
	}
	return p, nil
}

func (m *MockService) Delete(id int) (error) {
	if m.Error != "" {
		return fmt.Errorf(m.Error)
	}
	return nil
}
