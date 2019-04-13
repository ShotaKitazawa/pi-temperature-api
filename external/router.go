package external

import (
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/controllers"
	// "github.com/ShotaKitazawa/pi-temperature-api/external/mysql"
	"github.com/ShotaKitazawa/pi-temperature-api/external/rest"
	"github.com/ShotaKitazawa/pi-temperature-api/external/sqlite"

	"github.com/gin-gonic/gin"
)

// Router called by main.go
var Router *gin.Engine

func init() {
	r := gin.Default()
	v1 := r.Group("/api")

	logger := &Logger{}

	// db_conn := mysql.Connect()
	dbConn := sqlite.Connect()
	httpConn := rest.Request()

	InputController := controllers.NewInputController(dbConn, logger)
	OutputController := controllers.NewOutputController(httpConn, logger)

	v1.GET("/temperature", func(c *gin.Context) { InputController.Get1(c) })
	v1.POST("/aircon", func(c *gin.Context) { OutputController.Post(c) })

	Router = r
}
