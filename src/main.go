package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/route"
	"github.com/gin-gonic/gin"
)

func initialize() {
	config.InitOs()
	config.InitEnv()
}

func useMiddleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}

func main() {
	initialize()

	r := gin.New()

	useMiddleware(r)

	r.GET("/", controller.Ping)

	route.UseTodos(r)

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
