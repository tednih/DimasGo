// main.go
package main

import (
	"DimasGO/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Group routes under /tickets
	ticketGroup := r.Group("/tickets")
	{
		ticketGroup.GET("/", controllers.GetTickets)
		ticketGroup.POST("/", controllers.CreateTicket)
	}

	r.Run(":8080")
}
