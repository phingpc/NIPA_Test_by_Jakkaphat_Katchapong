package service

import (
	"NIPA_Test_Backend/internal/models"
	"NIPA_Test_Backend/internal/repository"
)

type TicketService interface {
	CreateTicket(ticket models.TicketInput) (models.Ticket, error)
	GetTicketAll() ([]models.Ticket, error)
	GetTicketByID(id int) (models.Ticket, error)
	UpdateTicketByID(id int, ticket models.TicketInput) (models.Ticket, error)
}

type ticketService struct {
	repo repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) TicketService {
	return &ticketService{repo: repo}
}

func (s *ticketService) CreateTicket(ticket models.TicketInput) (models.Ticket, error) {
	return s.repo.CreateTicket(ticket)
}

func (s *ticketService) GetTicketAll() ([]models.Ticket, error) {
	return s.repo.GetTicketAll()
}

func (s *ticketService) GetTicketByID(id int) (models.Ticket, error) {
	return s.repo.GetTicketByID(id)
}

func (s *ticketService) UpdateTicketByID(id int, ticket models.TicketInput) (models.Ticket, error) {
	return s.repo.UpdateTicketByID(id, ticket)
}