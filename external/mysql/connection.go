package mysql

import (
	"github.com/ShotaKitazawa/pi-temperature-api/adapter/gateway"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	url := "root:@tcp(db:3306)/pi-temterature"
	db, err = gorm.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&gateway.Input{}) {
		if err := db.Table("input").CreateTable(&gateway.Input{}).Error; err != nil {
			panic(err)
		}
	}

	return db
}

func CloseConn() {
	db.Close()
}
