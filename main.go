package main

import (
	"florist-gin/helpers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	userRepo "florist-gin/drivers/databases/users"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&userRepo.User{},
	)
}

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// dbPort := os.Getenv("DB_PORT")

	router := gin.Default()

	db, err := helpers.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	dbMigrate(db)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Vanilla Florist!")
	})

	port := ":8000"
	err = router.Run(port)
	if err != nil {
		log.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
