package main

import (
	"github.com/ShotaKitazawa/pi-temperature-api/external"
	// "github.com/ShotaKitazawa/pi-temperature-api/external/mysql"
	"github.com/ShotaKitazawa/pi-temperature-api/external/sqlite"
)

func main() {
	// defer mysql.CloseConn()
	defer sqlite.CloseConn()

	external.Router.Run("0.0.0.0:10080")
}
