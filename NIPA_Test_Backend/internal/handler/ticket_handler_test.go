package handler

import (
	"NIPA_Test_Backend/internal/models"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type MockTicketService struct {
	mock.Mock
}

func (m *MockTicketService) GetTicketByID(id int) (models.Ticket, error) {
	args := m.Called(id)
	return args.Get(0).(models.Ticket), args.Error(1)
}


func TestGetTicketByID(t *testing.T) {
	mockService := new(MockTicketService)

	// ตั้งค่า mock ให้ return ticket ที่ต้องการเมื่อเรียกฟังก์ชัน GetTicketByID
	mockService.On("GetTicketByID", 1).Return(models.Ticket{
		ID:          1,
		Title:       "Test Ticket",
		Description: "Test description",
		Status:      "pending",
		Contact:     "contact@example.com",
	}, nil)

	// เรียกฟังก์ชันจริง
	ticket, err := mockService.GetTicketByID(1)

	// ตรวจสอบผลลัพธ์
	assert.NoError(t, err)
	assert.Equal(t, 1, ticket.ID)
	assert.Equal(t, "Test Ticket", ticket.Title)
	mockService.AssertExpectations(t) // ตรวจสอบว่า mock ถูกเรียกตามที่คาด
}
