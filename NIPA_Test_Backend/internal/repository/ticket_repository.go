package repository

import (
	"database/sql"

	"NIPA_Test_Backend/internal/config"
	"NIPA_Test_Backend/internal/models"
	"fmt"

	_ "github.com/lib/pq"
)

type TicketRepository interface {
	CreateTicket(ticket models.TicketInput) (models.Ticket, error)
	GetTicketAll() ([]models.Ticket, error)
	GetTicketByID(id int) (models.Ticket, error)
	UpdateTicketByID(id int, ticket models.TicketInput) (models.Ticket, error)
}

type ticketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func ConnectDB(cfg config.Config) (*sql.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	return sql.Open("postgres", dsn)
}

func CheckDBConnection(db *sql.DB) error {
	return db.Ping()
}

func (r *ticketRepository) CreateTicket(ticket models.TicketInput) (models.Ticket, error) {
	query := `INSERT INTO tickets (title, description, contact, status) 
	          VALUES ($1, $2, $3, $4) 
	          RETURNING id, title, description, contact, status, created, updated`
	
	var newTicket models.Ticket
	err := r.db.QueryRow(query, ticket.Title, ticket.Description, ticket.Contact, ticket.Status).
		Scan(&newTicket.ID, &newTicket.Title, &newTicket.Description, &newTicket.Contact, &newTicket.Status, &newTicket.CreatedAt, &newTicket.UpdatedAt)
	
	if err != nil {
		return models.Ticket{}, err
	}
	return newTicket, nil
}


func (r *ticketRepository) GetTicketAll() ([]models.Ticket, error) {
	query := `SELECT id, title, description, contact, created, updated, status FROM tickets ORDER BY updated DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		err := rows.Scan(&ticket.ID, &ticket.Title, &ticket.Description, &ticket.Contact, &ticket.CreatedAt, &ticket.UpdatedAt, &ticket.Status)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (r *ticketRepository) GetTicketByID(id int) (models.Ticket, error) {
	query := `SELECT id, title, description, contact, created, updated, status FROM tickets WHERE id = $1`
	var ticket models.Ticket
	err := r.db.QueryRow(query, id).Scan(&ticket.ID, &ticket.Title, &ticket.Description, &ticket.Contact, &ticket.CreatedAt, &ticket.UpdatedAt, &ticket.Status)
	if err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (r *ticketRepository) UpdateTicketByID(id int, ticket models.TicketInput) (models.Ticket, error) {
    query := `UPDATE tickets SET title = $1, description = $2, status = $3, contact = $4 WHERE id = $5 RETURNING id, title, description, status, contact, created, updated`
    var updatedTicket models.Ticket
    err := r.db.QueryRow(query, ticket.Title, ticket.Description, ticket.Status, ticket.Contact, id).Scan(&updatedTicket.ID, &updatedTicket.Title, &updatedTicket.Description, &updatedTicket.Status, &updatedTicket.Contact, &updatedTicket.CreatedAt, &updatedTicket.UpdatedAt)
    if err != nil {
        return models.Ticket{}, err
    }
    return updatedTicket, nil
}
