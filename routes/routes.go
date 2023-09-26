// routes/routes.go
package routes

import (
	"DimasGo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	ticketGroup := r.Group("/tickets")
	{
		ticketGroup.GET("/", controllers.GetTickets)
<<<<<<< HEAD
		ticketGroup.POST("/", controllers.CreateTicket)
		ticketGroup.GET("/:id", controllers.GetTicketByID)
		ticketGroup.PUT("/:id", controllers.UpdateTicket)
		ticketGroup.DELETE("/:id", controllers.DeleteTicket)
=======
>>>>>>> 363d9fde51d208b3213bdc730f0d24fa744c325c
	}

	return r
}
