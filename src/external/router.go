package external

import (
	"github.com/gin-gonic/gin"
	"github.com/nakabonne/cleanarc-sample/src/adapter/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	handler := NewMysqlHandler()
	defer handler.CloseConn()
	userController := controllers.NewUserController(handler)
}
