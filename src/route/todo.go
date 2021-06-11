package route

import (
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/gin-gonic/gin"
)

func UseTodos(r *gin.Engine) {
	todos := r.Group("/todos")
	{
		todos.GET("/", controller.GetTodos)
		todos.POST("/", controller.AddTodo)
		todos.GET("/:todoId", controller.GetTodoById)
		todos.PATCH("/:todoId", controller.PatchTodoById)
		todos.DELETE("/:todoId", controller.DeleteTodoById)
	}
}
