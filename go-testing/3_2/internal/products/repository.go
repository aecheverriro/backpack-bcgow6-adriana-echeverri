package products

import (
	"clase3_parte2/internal/domain"
	"clase3_parte2/pkg/store"
	"fmt"
)

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, name, producType string, count int, price float64) (domain.Product, error)
	LastID() (int, error)
	Delete(id int) (error)
	Update(id int, name string, typeProduct string, count int, price float64) (domain.Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return ps, nil
}

func (r *repository) Store(id int, name, producType string, count int, price float64) (domain.Product, error) {
	var ps []domain.Product

	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailReading)
	}

	p := domain.Product{ID: id, Name: name, Type: producType, Count: count, Price: price}
	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, fmt.Errorf(FailWriting)
	}

	return p, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil
}

func (r *repository) Delete(id int) error {
	var ps []domain.Product
	deleted := false
	var index int

	if err := r.db.Read(&ps); err != nil {
		return fmt.Errorf(FailReading)
	}

	for i, value := range ps {
		if value.ID == id {
			deleted = true
			index = i
		}
	}
	if !deleted {
		return fmt.Errorf(ProductNotFound, id)
	}

	ps = append(ps[:index], ps[index + 1:]...)
	if err := r.db.Write(ps); err != nil {
		return fmt.Errorf(FailWriting)
	}

	return nil

}

func (r *repository) Update(id int, name string, typeProduct string, count int, price float64) (domain.Product, error) {
	var allProducts []domain.Product
	updatedProducts := domain.Product{
		ID: id, 
		Name: name, 
		Type: typeProduct,
		Count: count,
		Price: price,
	}
	
	if err := r.db.Read(&allProducts); err != nil {
		return domain.Product{}, err
	}

	for index, prod := range allProducts {
		if prod.ID == updatedProducts.ID {
			allProducts[index] = updatedProducts

			if err := r.db.Write(allProducts); err != nil {
				return domain.Product{}, err
			}
			return updatedProducts, nil
		}
	}

	return domain.Product{}, fmt.Errorf(ProductNotFound, id)
}