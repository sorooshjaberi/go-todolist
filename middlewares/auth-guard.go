package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todolist/constants"
	"todolist/lib/ginLib"
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

		username, err := encryptionUtils.ParseJWT(tokenString)

		if err != nil {
			context.JSON(http.StatusUnauthorized, ginLib.ResponseModel{
				Error: constants.ErrInvalidJWTToken.Error(),
			})
			context.Abort()
			errorsUtils.HandleErrorSoft(err)
			return
		}

		context.Set(constants.Keys.RequestUser, username)
		context.Next()

	}
}
