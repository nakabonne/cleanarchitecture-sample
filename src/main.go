package main

import (
	"github.com/nakabonne/cleanarc-sample/src/external"
	"github.com/nakabonne/cleanarc-sample/src/external/mysql"
)

func main() {
	defer mysql.CloseConn()

	external.Router.Run()
}
