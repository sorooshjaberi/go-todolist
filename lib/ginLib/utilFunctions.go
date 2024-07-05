package ginLib

import (
	"github.com/gin-gonic/gin"
	"todolist/constants"
	"todolist/models"
)

func GetUserFromContext(context *gin.Context) (*models.User, error) {
	currentUser, exists := context.Get(constants.Keys.RequestUser)

	if !exists {
		return nil, constants.ErrInternalServer
	}

	currentUserModel := currentUser.(models.User)

	return &currentUserModel, nil
}
