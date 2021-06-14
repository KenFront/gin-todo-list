package route

import (
	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/middleware"
	"github.com/KenFront/gin-todo-list/src/model"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func UseTodos(r *gin.Engine) {
	todos := r.Group("/todos")
	middleware.UseAuthGuard(todos)
	{
		todos.GET("/", controller.GetTodos)
		todos.POST("/", controller.AddTodo(model.AddTodoProps{
			Db:        config.GetDB(),
			GetUserId: util.GetUserId,
		}))
		todos.GET("/:todoId", controller.GetTodoById)
		todos.PATCH("/:todoId", controller.PatchTodoById)
		todos.DELETE("/:todoId", controller.DeleteTodoById)
	}
}
