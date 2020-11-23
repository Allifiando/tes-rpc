package routes

import (
	"github.com/gin-gonic/gin"
	controllers "santara.co.id/set/market-common-service-grpc/client/controller"
)

// Init ...
func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})
	userController := controllers.User{}
	r.GET("/bearer", userController.GetBearer)
	return r
}
