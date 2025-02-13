package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"NIPA_Test_Backend/internal/models"
	"NIPA_Test_Backend/internal/service"

	"NIPA_Test_Backend/internal/config"
	"NIPA_Test_Backend/internal/repository"
	"log"
)

type TicketHandler struct {
	TicketService service.TicketService
}

func NewTicketHandler(ts service.TicketService) *TicketHandler {
	return &TicketHandler{TicketService: ts}
}

func (h *TicketHandler) HealthCheck(c *gin.Context) {
	cfg := config.LoadConfig()
	db, err := repository.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()
	if err := repository.CheckDBConnection(db); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"detail": "Database connection failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "database": "connected"})
}

func (h *TicketHandler) CreateTicket(c *gin.Context) {
	var ticketInput models.TicketInput
	if err := c.ShouldBindJSON(&ticketInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err := h.TicketService.CreateTicket(ticketInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ticket)
}

func (h *TicketHandler) GetTicketAll(c *gin.Context) {
	tickets, err := h.TicketService.GetTicketAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) GetTicketByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err := h.TicketService.GetTicketByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) UpdateTicketByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ticketInput models.TicketInput
	if err := c.ShouldBindJSON(&ticketInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err := h.TicketService.UpdateTicketByID(id, ticketInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}