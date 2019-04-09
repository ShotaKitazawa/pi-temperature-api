package mysql

import (
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/gateway"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("sqlite", "hoge")

	if err != nil {
		panic(err)
	}
	if !db.HasTable(&gateway.User{}) {
		if err := db.Table("users").CreateTable(&gateway.User{}).Error; err != nil {
			panic(err)
		}
	}
	return db
}

func CloseConn() {
	db.Close()
}
