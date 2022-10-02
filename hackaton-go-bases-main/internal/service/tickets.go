package service

import (
	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

// Un método para crear un nuevo ticket añadir al registro.
func (b *bookings) Create(t Ticket) (Ticket, error) {
	err := file.Write(t)
	return t, err
}

func (b *bookings) Read(id int) (Ticket, error) {
	data, err := file.Read()

	for _, ticket := range data {
		if ticket.Id == id {
			return ticket, err
		}
	}
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
