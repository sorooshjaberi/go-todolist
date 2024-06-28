package auth

import (
	"booking/constants"
	"booking/models"
	usersService "booking/services/users"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitiateRouter(server *gin.RouterGroup) {
	authRouter := server.Group("/auth")

	authRouter.POST("/login", loginHandler)
	authRouter.POST("/signup", signupHandler)
}

func loginHandler(context *gin.Context) {
	var body models.User

	//cast body
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//handle password verification
	claimedUser, err := usersService.Login(body.Username, body.Password)

	if err != nil {
		if errors.Is(err, constants.ErrEmptyCredentials) || errors.Is(err, constants.ErrInvalidCredentials) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if errors.Is(err, constants.ErrUserNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": claimedUser})
	//context.Set(constants.RESPONSE_JSON, "hello")
}

func signupHandler(context *gin.Context) {
	var body models.User

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newUser, err := usersService.Signup(body.Username, body.Password)

	if err != nil {
		if errors.Is(err, constants.ErrEmptyCredentials) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": newUser})
}
