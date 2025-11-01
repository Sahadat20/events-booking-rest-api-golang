package routes

import (
	"github.com/Sahadat20/events-booking-rest-api-golang/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authentiated := server.Group("/")
	authentiated.Use(middlewares.Authenticate) // always run the midleware int all authentiated group routes
	authentiated.POST("/events", createEvent)
	authentiated.PUT("/events/:id", updateEvent)
	authentiated.DELETE("/events/:id", deleteEvent)
	authentiated.POST("/events/:id/register", registerForEvent)
	authentiated.DELETE("/events/:id/register", cancenRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
