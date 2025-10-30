package main

import (
	// "github.com/Sahadat20/events-booking-rest-api-golang/db"
	"github.com/Sahadat20/events-booking-rest-api-golang/db"
	"github.com/Sahadat20/events-booking-rest-api-golang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080

}
