package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// IndexGet handler for test
// Method: GET
// Path:   `/`
func IndexGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "pong",
		"time": time.Now().Local(),
		"ip":   c.ClientIP(),
	})
}
