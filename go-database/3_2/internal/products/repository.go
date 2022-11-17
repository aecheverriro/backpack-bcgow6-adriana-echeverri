package products

import (
	"3_2/internal/domain"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Repository interface {
	Store(p domain.Product) (int, error)
	GetByName(name string) (domain.Product, error)
	GetAll(id int) ([]domain.Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByName(name string) (domain.Product, error) {
	stmt, err := r.db.Prepare(GET_BY_NAME_QUERY)
	if err != nil {
		return domain.Product{}, errors.New("error al preparar la consulta")
	}
	defer stmt.Close()

	var product domain.Product
	err = stmt.QueryRow(name).Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return domain.Product{}, errors.New("error registrando")
	}

	return product, nil
}

func (r *repository) Store(p domain.Product) (int, error) {
	stmt, err := r.db.Prepare(INSERT_QUERY)
	if err != nil {
		return 0, fmt.Errorf("error al preparar la consulta - error %s", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Type, p.Count, p.Price, p.ID_wh)
	if err != nil {
		return 0, fmt.Errorf("error al hacer la consulta - error %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("error al obtener el ultimo id")
	}

	return int(id), nil
}

func (r *repository) GetAll(id int) ([]domain.Product, error) {
	db := r.db
		
	rows, err := db.Query(GET_BY_ID_QUERY, id)
	if err != nil {
		return nil, err
	}
	var products []domain.Product

	for rows.Next() {
		var p domain.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price, &p.ID_wh)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *repository) Delete(id int) error {
	db := r.db
	stmt, err := db.Prepare(DELETE_QUERY)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	
	if err != nil {
		return err
	}

	return nil
}
