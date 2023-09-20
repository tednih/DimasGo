// controllers/ticket_controller.go
package controllers

import (
	"DimasGO/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var tickets []models.Ticket

// Handler untuk endpoint GET /tickets
func GetTickets(c *gin.Context) {
	c.JSON(http.StatusOK, tickets)
}

// Handler untuk endpoint POST /tickets
func CreateTicket(c *gin.Context) {
	var newTicket models.Ticket
	if err := c.ShouldBindJSON(&newTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tickets = append(tickets, newTicket)
	c.JSON(http.StatusCreated, newTicket)
}
