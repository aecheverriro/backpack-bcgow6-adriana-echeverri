package tickets

import (
	"context"
	"fmt"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/desafio-goweb-adrianaecheverri/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTicketsByCountry(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetTicketsByCountry(ctx context.Context, destination string) (int, error){
	var amount int

	if len(r.db) == 0 {
		return 0, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			amount++
		}
	}

	return amount, nil
}

func (r *repository) AverageDestination(ctx context.Context, destination string) (float64, error){
	var amount float64

	if len(r.db) == 0 {
		return 0, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			amount++
		}
	}

	return amount, nil
}
