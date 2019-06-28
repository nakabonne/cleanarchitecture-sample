package external

import (
	"cleanarchitecture-sample/adapter/controllers"
	"cleanarchitecture-sample/external/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	logger := &Logger{}
	conn := mysql.Connect()
	userController := controllers.NewUserController(conn, logger)

	router.GET("/hc", func(c *gin.Context) { c.String(http.StatusOK, "hello") })
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })

	Router = router
}
