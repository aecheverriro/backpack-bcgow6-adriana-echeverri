package tickets

import (
	"context"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/desafio-goweb-adrianaecheverri/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTicketsByCountry(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	ticks, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return ticks, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	ticketsDest, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}

	return ticketsDest, nil
}

func (s *service) GetTicketsByCountry(ctx context.Context, destination string) (int, error) {

	numberTickets, err := s.repository.GetTicketsByCountry(ctx, destination)
	if err != nil {
		return 0, err
	}

	return numberTickets, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {

	averageTickets, err := s.repository.AverageDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	return averageTickets, nil
}

