// controllers/ticket_controller.go
package controllers

import (
	"DimasGo/config"
	"DimasGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
