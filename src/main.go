package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/middleware"
	"github.com/KenFront/gin-todo-list/src/route"
	"github.com/gin-gonic/gin"
)

func initialize() {
	config.InitOs()
	config.InitEnv()
}

func useGlobalMiddlewares(r *gin.Engine) {
	middleware.UseLogger(r)
	middleware.UseRecovery(r)
}

func main() {
	initialize()

	r := gin.New()

	useGlobalMiddlewares(r)

	r.GET("/", controller.Ping)

	route.UseTodos(r)

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
