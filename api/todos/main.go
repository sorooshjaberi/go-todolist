package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolist/constants"
	"todolist/lib/ginLib"
	"todolist/models"
	"todolist/services/todos"
)

func RegisterRouter(server *gin.RouterGroup) {
	todosRouter := server.Group("/todos")
	todosRouter.GET("/", getAllTodosHandler)
	todosRouter.POST("/", createTodoHandler)
	todosRouter.PUT("/:todoId", editTodoHandler)
	todosRouter.DELETE("/:todoId", deleteTodoHandler)
}

func getAllTodosHandler(context *gin.Context) {
	page, pageErr := strconv.Atoi(context.DefaultQuery("page", "1"))
	if pageErr != nil {
		context.JSON(http.StatusBadRequest, ginLib.ResponseModel{Error: pageErr.Error()})
		return
	}

	perPage, perPageErr := strconv.Atoi(context.DefaultQuery("per_page", "10"))
	if perPageErr != nil {
		context.JSON(http.StatusBadRequest, ginLib.ResponseModel{Error: perPageErr.Error()})
		return
	}

	currentUser, getUserError := ginLib.GetUserFromContext(context)

	if getUserError != nil {
		context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: constants.ErrInternalServer.Error()})
		return
	}

	currentUserId := currentUser.ID

	allTodos, allTodosErr := todos.GetAllTodos(currentUserId, page, perPage)

	if allTodosErr != nil {
		context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: allTodosErr.Error()})
		return
	}

	context.JSON(200, ginLib.ResponseModel{
		Data: allTodos,
	})
}

func createTodoHandler(context *gin.Context) {
	currentUser, getUserError := ginLib.GetUserFromContext(context)
	if getUserError != nil {
		context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: constants.ErrInternalServer.Error()})
		return
	}

	var newTodo models.Todo

	if err := context.ShouldBindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, ginLib.ResponseModel{
			Error: err.Error(),
		})
		return
	}

	newTodo.UserID = currentUser.ID

	createdTodo, createTodoError := todos.CreateTodo(newTodo)

	if createTodoError != nil {
		context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: createTodoError.Error()})
		return
	}

	context.JSON(http.StatusCreated, ginLib.ResponseModel{Data: createdTodo})
}

func editTodoHandler(context *gin.Context) {

	todoId, convertTodoIdToStringError := strconv.Atoi(context.Param("todoId"))

	if convertTodoIdToStringError != nil {
		context.JSON(http.StatusBadRequest, ginLib.ResponseModel{Error: constants.ErrIdShouldBeNumber.Error()})
		return
	}

	currentUser, getUserError := ginLib.GetUserFromContext(context)

	if getUserError != nil {
		context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: constants.ErrInternalServer.Error()})
		return
	}

	var todo models.Todo

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, ginLib.ResponseModel{
			Error: err.Error(),
		})
		return
	}

	todo.ID = uint(todoId)
	todo.UserID = currentUser.ID

	updatedTodo, editTodoError := todos.EditTodo(todo)

	if editTodoError != nil {
		context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: editTodoError.Error()})
		return
	}

	context.JSON(http.StatusCreated, ginLib.ResponseModel{Data: updatedTodo})
}

func deleteTodoHandler(context *gin.Context) {
	todoId, convertTodoIdToIntError := strconv.Atoi(context.Param("todoId"))

	if convertTodoIdToIntError != nil {
		context.JSON(http.StatusBadRequest, ginLib.ResponseModel{Error: constants.ErrIdShouldBeNumber.Error()})
		return
	}
	currentUser, getUserError := ginLib.GetUserFromContext(context)

	if getUserError != nil {
		context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: constants.ErrInternalServer.Error()})
		return
	}

	var todo models.Todo

	todo.UserID = currentUser.ID

	todo.ID = uint(todoId)

	deletedTodo, deleteTodoError := todos.DeleteTodo(todo)

	if deleteTodoError != nil {
		context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: deleteTodoError.Error()})
		return
	}

	context.JSON(http.StatusNoContent, ginLib.ResponseModel{Data: deletedTodo})

}
