package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/constants"
	"todolist/lib/ginLib"
)

func RegisterRouter(server *gin.RouterGroup) {
	todosRouter := server.Group("/todos")
	todosRouter.GET("/", func(context *gin.Context) {
		currentUser, exists := context.Get(constants.Keys.RequestUser)

		if !exists {
			context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{Error: constants.ErrInternalServer.Error()})
			return
		}

		context.JSON(200, ginLib.ResponseModel{
			Data: currentUser,
		})
	})
}
