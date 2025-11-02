package main

import (
	// "github.com/Sahadat20/events-booking-rest-api-golang/db"
	"log"
	"os"

	"github.com/Sahadat20/events-booking-rest-api-golang/db"
	"github.com/Sahadat20/events-booking-rest-api-golang/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":" + port) //localhost:8080

}
