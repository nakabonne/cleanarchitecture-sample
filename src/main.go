package main

import (
	"external"
	"external/mysql"
)

func main() {
	defer mysql.CloseConn()

	external.Router.Run()
}
