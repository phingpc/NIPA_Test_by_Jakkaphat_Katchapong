package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"NIPA_Test_Backend/internal/config"
	"NIPA_Test_Backend/internal/handler"
	"NIPA_Test_Backend/internal/repository"
	"NIPA_Test_Backend/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	db, err := repository.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	ticketRepo := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepo)
	ticketHandler := handler.NewTicketHandler(ticketService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/health", ticketHandler.HealthCheck)

	authRequired := r.Group("/api/v1")
	{
		authRequired.POST("/ticket", ticketHandler.CreateTicket)
		authRequired.GET("/ticket", ticketHandler.GetTicketAll)
		authRequired.GET("/ticket/:id", ticketHandler.GetTicketByID)
		authRequired.PUT("/ticket/:id", ticketHandler.UpdateTicketByID)
	}

	port := cfg.APIPORT
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
