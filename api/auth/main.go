package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/constants"
	"todolist/lib/ginLib"
	"todolist/models"
	usersService "todolist/services/users"
)

func RegisterRouter(server *gin.RouterGroup) {
	authRouter := server.Group("/auth")

	authRouter.POST("/login", loginHandler)
	authRouter.POST("/signup", signupHandler)
}

func loginHandler(context *gin.Context) {
	var body models.User

	//cast body
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, ginLib.ResponseModel{
			Error: err.Error(),
		})
		return
	}

	//handle password verification
	jwtToken, err := usersService.Login(body.Username, body.Password)

	if err != nil {
		if errors.Is(err, constants.ErrEmptyCredentials) || errors.Is(err, constants.ErrInvalidCredentials) {
			context.JSON(http.StatusBadRequest, ginLib.ResponseModel{
				Error: err.Error(),
			})
		} else if errors.Is(err, constants.ErrUserNotFound) {
			context.JSON(http.StatusNotFound, ginLib.ResponseModel{
				Error: err.Error(),
			})
		} else {
			context.JSON(http.StatusInternalServerError, ginLib.ResponseModel{
				Error: err.Error(),
			})
		}
		return
	}

	context.JSON(http.StatusOK, ginLib.ResponseModel{Data: gin.H{"token": jwtToken}})
}

func signupHandler(context *gin.Context) {
	var body models.User

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, ginLib.ResponseModel{
			Error: err.Error(),
		})
	}

	_, err := usersService.Signup(body.Username, body.Password)

	if err != nil {
		if errors.Is(err, constants.ErrEmptyCredentials) {
			context.JSON(http.StatusBadRequest, ginLib.ResponseModel{
				Error: err.Error(),
			})
		} else {
			context.JSON(http.StatusInternalServerError,
				ginLib.ResponseModel{
					Error: err.Error(),
				},
			)
		}
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
