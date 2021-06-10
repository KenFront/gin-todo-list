package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var DB *gorm.DB

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LinkDb() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv(("POSTGRES_PORT_HOST")))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

}

func Ping(c *gin.Context) {
	c.Abort()
}

func GetTodos(c *gin.Context) {
	var todos []Todo

	result := DB.Find(&todos)

	if result.Error == nil {

		c.JSON(http.StatusOK, gin.H{
			"data": todos,
		})

	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": result.Error,
		})
	}
}

func main() {
	GetEnv()
	LinkDb()
	r := gin.Default()
	r.GET("/", Ping)
	r.GET("/todos", GetTodos)

	r.Run()
}
