// controllers/ticket_controller.go
package controllers

import (
	"DimasGo/config"
	"DimasGo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
func CreateTicket(c *gin.Context) {
	var newTicket models.Ticket

	// Bind data JSON yang dikirim dalam permintaan ke struct newTicket
	if err := c.ShouldBindJSON(&newTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
=======
func GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	rows, err := config.DB.Query("SELECT * FROM tickets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
>>>>>>> 363d9fde51d208b3213bdc730f0d24fa744c325c
		return
	}
	defer rows.Close()

<<<<<<< HEAD
	// Simpan tiket baru ke dalam database
	_, err := config.DB.Exec("INSERT INTO tickets (movie, price) VALUES (?, ?)", newTicket.Movie, newTicket.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTicket)
=======
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(&ticket.ID, &ticket.Movie, &ticket.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tickets = append(tickets, ticket)
	}

	c.JSON(http.StatusOK, tickets)
>>>>>>> 363d9fde51d208b3213bdc730f0d24fa744c325c
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
	// Ambil ID tiket dari URL parameter
	idParam := c.Param("id")

	// Konversi ID dari string ke int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	// Query database untuk mendapatkan tiket berdasarkan ID
	var ticket models.Ticket
	err = config.DB.QueryRow("SELECT * FROM tickets WHERE id=?", id).Scan(&ticket.ID, &ticket.Movie, &ticket.Price)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func UpdateTicket(c *gin.Context) {
	// Ambil ID tiket dari URL parameter
	idParam := c.Param("id")

	// Konversi ID dari string ke int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	// Bind data JSON yang dikirim dalam permintaan ke struct ticket
	var updatedTicket models.Ticket
	if err := c.ShouldBindJSON(&updatedTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update tiket dalam database
	_, err = config.DB.Exec("UPDATE tickets SET movie=?, price=? WHERE id=?", updatedTicket.Movie, updatedTicket.Price, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully"})
}

func DeleteTicket(c *gin.Context) {
	// Ambil ID tiket dari URL parameter
	idParam := c.Param("id")

	// Konversi ID dari string ke int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	// Hapus tiket dari database berdasarkan ID
	_, err = config.DB.Exec("DELETE FROM tickets WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
