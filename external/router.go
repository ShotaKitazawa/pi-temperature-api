package external

import (
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/controllers"
	// "github.com/ShotaKitazawa/pi-temperature-api/external/mysql"
	"github.com/ShotaKitazawa/pi-temperature-api/external/rest"
	"github.com/ShotaKitazawa/pi-temperature-api/external/sqlite"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router called by main.go
var Router *gin.Engine

func init() {
	r := gin.Default()
	r.Use(cors.Default())

	v := r.Group("/api")

	logger := &Logger{}

	// dbConn := mysql.Connect(mysql.GetEnv())
	dbConn := sqlite.Connect(sqlite.GetEnv())
	httpConn := rest.Request(rest.GetEnv())

	InputController := controllers.NewInputController(dbConn, logger)
	OutputController := controllers.NewOutputController(dbConn, httpConn, logger)

	v.GET("/temperature", func(c *gin.Context) { InputController.Get1(c) })
	v.POST("/aircon", func(c *gin.Context) { OutputController.Post(c) })

	Router = r
}
