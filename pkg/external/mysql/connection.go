package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/nakabonne/cleanarchitecture-sample/pkg/adapter/gateway"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:@tcp(db:3306)/hoge")

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
