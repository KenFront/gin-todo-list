package route

import (
	"github.com/KenFront/gin-todo-list/src/config"
	"github.com/KenFront/gin-todo-list/src/controller"
	"github.com/KenFront/gin-todo-list/src/middleware"
	"github.com/KenFront/gin-todo-list/src/util"
	"github.com/gin-gonic/gin"
)

func UseTodos(r *gin.Engine) {
	todos := r.Group("/todos")
	db := config.GetDB()
	middleware.UseAuthGuard(todos)
	{
		todos.GET("/", controller.GetTodos(controller.GetTodosProps{
			Db:               db,
			GetUserIdByToken: util.GetUserIdByToken,
		}))
		todos.POST("/", controller.AddTodo(controller.AddTodoProps{
			Db:               db,
			GetUserIdByToken: util.GetUserIdByToken,
			GetNewTodoId:     util.GetNewUserId,
		}))
		todos.GET("/:todoId", controller.GetTodoById(controller.GetTodoByIdProps{
			Db:               db,
			GetUserIdByToken: util.GetUserIdByToken,
		}))
		todos.PATCH("/:todoId", controller.PatchTodoById)
		todos.DELETE("/:todoId", controller.DeleteTodoById(controller.DeleteTodoProps{
			Db:               db,
			GetUserIdByToken: util.GetUserIdByToken,
		}))
	}
}
