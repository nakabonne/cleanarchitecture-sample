package mysql

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	if isDevelop() {
		db, err = gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION_DEV"))
	} else {
		db, err = gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	}

	if err != nil {
		panic(err.Error)
	}
	return db
}

func isDevelop() bool {
	return os.Getenv("DEV") == "1"
}

func CloseConn() {
	db.Close()
}
