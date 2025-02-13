package service

import (
	"NIPA_Test_Backend/internal/models"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTicketRepository struct {
	mock.Mock
}

func (m *MockTicketRepository) CreateTicket(ticket models.TicketInput) (models.Ticket, error) {
	args := m.Called(ticket)
	return args.Get(0).(models.Ticket), args.Error(1)
}

func (m *MockTicketRepository) GetTicketAll() ([]models.Ticket, error) {
	args := m.Called()
	return args.Get(0).([]models.Ticket), args.Error(1)
}

func (m *MockTicketRepository) GetTicketByID(id int) (models.Ticket, error) {
	args := m.Called(id)
	return args.Get(0).(models.Ticket), args.Error(1)
}

func (m *MockTicketRepository) UpdateTicketByID(id int, ticket models.TicketInput) (models.Ticket, error) {
	args := m.Called(id, ticket)
	return args.Get(0).(models.Ticket), args.Error(1)
}

func TestCreateTicket(t *testing.T) {
	mockRepo := new(MockTicketRepository)
	service := NewTicketService(mockRepo)

	ticketInput := models.TicketInput{
		Title:       "Mocked Ticket",
		Description: "This is a mocked ticket",
		Status:      "open",
		Contact:     "mock@example.com",
	}

	mockRepo.On("CreateTicket", ticketInput).Return(models.Ticket{
		ID:          1,
		Title:       ticketInput.Title,
		Description: ticketInput.Description,
		Contact:     ticketInput.Contact,
		Status:      ticketInput.Status,
	}, nil)

	ticket, err := service.CreateTicket(ticketInput)

	assert.Nil(t, err)
	assert.Equal(t, "Mocked Ticket", ticket.Title)
	mockRepo.AssertExpectations(t)
}
