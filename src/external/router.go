package external

import (
	"github.com/gin-gonic/gin"
	"github.com/nakabonne/cleanarc-sample/src/adapter/controllers"
	"github.com/nakabonne/cleanarc-sample/src/external/mysql"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	logger := &Logger{}

	conn := mysql.Connect()

	userController := controllers.NewUserController(conn, logger)

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })

	Router = router
}
