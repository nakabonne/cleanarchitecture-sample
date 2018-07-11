package mysql

import (
	"adapter/gateway"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:@tcp(db:3306)/hoge")

	if err != nil {
		panic(err)
	}
	db.Table("users").CreateTable(&gateway.User{})
	return db
}

func CloseConn() {
	db.Close()
}
