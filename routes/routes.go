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
	}

	return r
}
