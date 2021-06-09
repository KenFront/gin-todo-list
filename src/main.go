package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID          string `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func getEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	getEnv()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv(("POSTGRES_PORT_HOST")))

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			var todos []Todo

			db.Find(&todos)

			c.JSON(200, gin.H{
				"data": todos,
			})
		} else {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		}
	})
	r.Run()
}
