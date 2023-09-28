// controllers/ticket_controller.go
package controllers

import (
	"DimasGo/config"
	"DimasGo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTicket(c *gin.Context) {
	var newTicket models.Ticket

	if err := c.ShouldBindJSON(&newTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := config.DB.Exec("INSERT INTO tickets (movie, price) VALUES (?, ?)", newTicket.Movie, newTicket.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTicket)
}

func GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	rows, err := config.DB.Query("SELECT * FROM tickets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(&ticket.ID, &ticket.Movie, &ticket.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tickets = append(tickets, ticket)
	}

	c.JSON(http.StatusOK, tickets)
}

func GetTicketByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	var ticket models.Ticket
	err = config.DB.QueryRow("SELECT * FROM tickets WHERE id=?", id).Scan(&ticket.ID, &ticket.Movie, &ticket.Price)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func UpdateTicket(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	var updatedTicket models.Ticket
	if err := c.ShouldBindJSON(&updatedTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = config.DB.Exec("UPDATE tickets SET movie=?, price=? WHERE id=?", updatedTicket.Movie, updatedTicket.Price, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully"})
}

func DeleteTicket(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM tickets WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
