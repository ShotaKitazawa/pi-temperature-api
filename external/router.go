package external

import (
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/controllers"
	// "github.com/ShotaKitazawa/pi-temperature-api/external/mysql"
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

	InputController := controllers.NewInputController(dbConn, logger)
	// OutputController := controllers.NewOutputController(dbConn, logger)

	v1.GET("/temperature", func(c *gin.Context) { InputController.Get1(c) })

	Router = r
}
