package main

import (
	"github.com/nakabonne/cleanarchitecture-sample/pkg/external"
	"github.com/nakabonne/cleanarchitecture-sample/pkg/external/mysql"
)

func main() {
	defer mysql.CloseConn()

	external.Router.Run()
}
