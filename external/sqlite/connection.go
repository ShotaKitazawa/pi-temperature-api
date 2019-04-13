package sqlite

import (
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/gateway"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	url := "pi-temperature.sqlite3"
	db, err = gorm.Open("sqlite3", url)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&gateway.Input{}) {
		if err := db.Table("inputs").CreateTable(&gateway.Input{}).Error; err != nil {
			panic(err)
		}
	}

	return db
}

func CloseConn() {
	db.Close()
}
