package main

import (
	"cleanarchitecture-sample/external"
	"cleanarchitecture-sample/external/mysql"
)

func main() {
	defer mysql.CloseConn()

	external.Router.Run()
}
