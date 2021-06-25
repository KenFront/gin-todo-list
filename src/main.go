package main

import (
	"log"

	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/middleware"
	"github.com/KenFront/gin-todo-list/src/route"
	"github.com/gin-gonic/gin"
)

func useGlobalMiddlewares(r *gin.Engine) {
	middleware.UseCustomLogger(r)
	middleware.UseLogger(r)
	middleware.UseErrorHandler(r)
}

func main() {
	config.InitOs()

	r := gin.New()

	useGlobalMiddlewares(r)

	r.GET("/", controller.Ping)

	route.UseAuth(r)
	route.UseUser(r)
	route.UseTodos(r)

	if err := r.Run(":80"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
