package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;" json:"id"`
	Title       string    `gorm:"type:string" json:"title"`
	Description string    `gorm:"type:string" json:"description"`
	Status      string    `gorm:"type:string;" json:"status"`
	CreatedAt   time.Time `json:"createAt"`
	UpdatedAt   time.Time `json:"updateAt"`
}

type AddTodo struct {
	Title       string
	Description string
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

func AddTodos(c *gin.Context) {
	var payload AddTodo

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	todo := Todo{Title: payload.Title, Description: payload.Description}

	createActionResult := DB.Select("ID", "Title", "Description").Create(&todo)
	createdDataResult := DB.First(&todo)

	switch {
	case createActionResult.Error != nil:
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": createActionResult.Error,
		})
	case createdDataResult.Error != nil:
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": createdDataResult.Error,
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
		})
	}
}

func main() {
	GetEnv()
	LinkDb()
	r := gin.Default()
	r.GET("/", Ping)
	r.GET("/todos", GetTodos)
	r.POST("/todos", AddTodos)

	r.Run()
}
