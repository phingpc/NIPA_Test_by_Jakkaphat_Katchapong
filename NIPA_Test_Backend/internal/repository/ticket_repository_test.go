package repository

import (
	"NIPA_Test_Backend/internal/models"
	"NIPA_Test_Backend/internal/config"
	"testing"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateTicket(t *testing.T) {
	cfg := config.LoadConfig()
	db, err := ConnectDB(cfg)
	if err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	ticketRepo := NewTicketRepository(db)

	ticketInput := models.TicketInput{
		Title:       "Test Ticket",
		Description: "Test description",
		Status:      "pending",
		Contact:     "contact@example.com",
	}	

	ticket, err := ticketRepo.CreateTicket(ticketInput)
	if err != nil {
		t.Fatalf("Failed to create ticket: %v", err)
	}

	assert.Equal(t, ticketInput.Title, ticket.Title)
	assert.Equal(t, ticketInput.Description, ticket.Description)
	assert.Equal(t, ticketInput.Status, ticket.Status)
}

func TestGetTicketByID(t *testing.T) {
	cfg := config.LoadConfig()
	db, err := ConnectDB(cfg)
	if err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	ticketRepo := NewTicketRepository(db)

	// Create a ticket first
	ticketInput := models.TicketInput{
		Title:       "Test Ticket",
		Description: "Test description",
		Status:      "pending",
		Contact:     "contact@example.com",
	}
	

	ticket, err := ticketRepo.CreateTicket(ticketInput)
	if err != nil {
		t.Fatalf("Failed to create ticket: %v", err)
	}

	// Now test GetTicketByID
	retrievedTicket, err := ticketRepo.GetTicketByID(ticket.ID)
	if err != nil {
		t.Fatalf("Failed to get ticket by ID: %v", err)
	}

	assert.Equal(t, ticket.ID, retrievedTicket.ID)
}
