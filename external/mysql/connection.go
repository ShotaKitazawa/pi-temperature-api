package mysql

import (
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/gateway"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:@tcp(db:3306)/pi-temterature")

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&gateway.Input{}) {
		if err := db.Table("input").CreateTable(&gateway.Input{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&gateway.Output{}) {
		if err := db.Table("output").CreateTable(&gateway.Output{}).Error; err != nil {
			panic(err)
		}
	}

	return db
}

func CloseConn() {
	db.Close()
}
