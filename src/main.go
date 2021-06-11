package main

import (
	"fmt"
	"os"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitOs()
	config.InitEnv()

	r := gin.Default()

	r.GET("/", controller.Ping)
	r.GET("/todos", controller.GetTodos)
	r.POST("/todos", controller.AddTodo)
	r.GET("/todos/:todoId", controller.GetTodoById)
	r.PATCH("/todos/:todoId", controller.PatchTodoById)
	r.DELETE("/todos/:todoId", controller.DeleteTodoById)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
