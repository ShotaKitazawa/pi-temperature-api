package main

import (
	"github.com/ShotaKitazawa/pi-temperature-api/external"
	"github.com/ShotaKitazawa/pi-temperature-api/external/mysql"
)

func main() {
	defer mysql.CloseConn()

	external.Router.Run()
}
