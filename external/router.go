package external

import (
	"github.com/ShotaKitazawa/pi-temterature-api/adapter/controllers"
	"github.com/ShotaKitazawa/pi-temterature-api/external/mysql"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	r := gin.Default()
	v1 := r.Group("/api")

	logger := &Logger{}

	conn := mysql.Connect()

	userController := controllers.NewUserController(conn, logger)

	v1.POST("/users", func(c *gin.Context) { userController.Create(c) })

	Router = router
}

