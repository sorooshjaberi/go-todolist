package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"todolist/constants"
	"todolist/lib/ginLib"
	usersService "todolist/services/users"
	"todolist/utils/encryptionUtils"
	"todolist/utils/errorsUtils"
)

func AuthGuard() gin.HandlerFunc {
	return func(context *gin.Context) {

		tokenString := context.GetHeader("Authorization")

		if tokenString == "" || strings.Trim(tokenString, " ") == "Bearer" {
			context.JSON(http.StatusUnauthorized, ginLib.ResponseModel{
				Error: constants.ErrMissingToken.Error(),
			})
		}

		tokenString = tokenString[len("Bearer "):]

		userId, err := encryptionUtils.ParseJWT(tokenString)

		fmt.Println("userId", userId)

		if err != nil {
			context.JSON(http.StatusUnauthorized, ginLib.ResponseModel{
				Error: constants.ErrInvalidJWTToken.Error(),
			})
			context.Abort()
			errorsUtils.HandleErrorSoft(err)
			return
		}

		userIdUint, err := strconv.ParseUint(userId, 10, 64)

		errorsUtils.HandleErrorSoft(err)

		user, err := usersService.FindUserById(uint(userIdUint))

		if err != nil {
			context.JSON(http.StatusNotFound, ginLib.ResponseModel{
				Error: err.Error(),
			})
		}

		context.Set(constants.Keys.RequestUser, user)
		context.Next()

	}
}
