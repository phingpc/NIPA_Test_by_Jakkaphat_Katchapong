
package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketModel(t *testing.T) {
	ticket := Ticket{
		Title:       "Test Ticket",
		Description: "Test Description",
		Status:      "pending",
		Contact:     "contact@example.com",
	}

	
	assert.Equal(t, "Test Ticket", ticket.Title)
	assert.Equal(t, "Test Description", ticket.Description)
	assert.Equal(t, "pending", ticket.Status)
	assert.Equal(t, "contact@example.com", ticket.Contact)
}
