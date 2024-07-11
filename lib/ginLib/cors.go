package ginLib

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(server *gin.Engine) {

	corsConfig := cors.DefaultConfig()
	//corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowAllOrigins = true

	server.Use(cors.New(corsConfig))

}
