package main

import (
	"florist-gin/helpers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPort := os.Getenv("DB_PORT")

	// Initialize the Gin router
	router := gin.Default()

	// Initialize the database connection using environment variables
	db, err := helpers.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection works
	pingErr := db.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Vanilla Florist!")
	})

	// Start the HTTP server
	port := ":" + dbPort
	err = router.Run(port)
	if err != nil {
		log.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
