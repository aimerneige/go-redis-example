package router

import (
	"go-redis-example/internal/app/api/handler"
	"go-redis-example/internal/app/api/handler/data"

	"github.com/gin-gonic/gin"
)

// Setup setup api route
func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/", handler.IndexGet)
	r.GET("/data", data.Get)
	r.PUT("/data", data.Put)

	return r
}
