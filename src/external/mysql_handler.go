package external

import (
	"database/sql"

	"github.com/nakabonne/cleanarc-sample/src/adapter/interfaces"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlHandler struct {
	Conn *gorm.DB
}

func NewMysqlHandler() interfaces.DBHandler {
	var (
		handler *MysqlHandler
		db      *gorm.DB
		err     error
	)

	if isDevelop() {
		db, err = gorm.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASS")+"@/"+os.Getenv("GENEPSE_DBNAME")+"?charset="+os.Getenv("MYSQL_CHARSET")+"&parseTime="+os.Getenv("MYSQL_PARSETIME")+"&loc="+os.Getenv("MYSQL_LOC"))
	} else {
		db, err = gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	}

	if err != nil {
		panic(err.Error)
	}

	handler.Conn = db
	return handler
}

func isDevelop() bool {
	return os.Getenv("DEV") == "1"
}

func (h *MysqlHandler) CloseConn() {
	h.Conn.Close()
}
