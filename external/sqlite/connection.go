package sqlite

import (
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/gateway"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("sqlite3", "pi-temperature.sqlite3")

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&gateway.Input{}) {
		if err := db.Table("inputs").CreateTable(&gateway.Input{}).Error; err != nil {
			panic(err)
		}
	}
	//if !db.HasTable(&gateway.Output{}) {
	//	if err := db.Table("outputs").CreateTable(&gateway.Output{}).Error; err != nil {
	//		panic(err)
	//	}
	//}

	return db
}

func CloseConn() {
	db.Close()
}
