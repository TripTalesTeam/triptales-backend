package route

import (
	"github.com/gin-gonic/gin"
	"github.com/breezjirasak/triptales/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.CreateUser)

	return r
}