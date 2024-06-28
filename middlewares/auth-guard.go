package middlewares

import (
	"booking/constants"
	"booking/lib/ginLib"
	"booking/utils/encryptionUtil"
	"booking/utils/errorsUtils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthGuard() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, ginLib.ResponseModel{
				Error: constants.ErrMissingToken,
			})
		}

		tokenString = tokenString[len("Bearer "):]

		username, err := encryptionUtil.ParseJWT(tokenString)

		if err != nil {
			errorsUtils.HandleErrorSoft(err)
			return
		}

		context.Set(constants.Keys.RequestUser, username)

	}
}
