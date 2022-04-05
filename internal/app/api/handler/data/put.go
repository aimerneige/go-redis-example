package data

import (
	"go-redis-example/internal/app/api/cache"
	"go-redis-example/internal/app/api/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Put handler for update data
// Method: PUT
// Path:   `/data/`
func Put(c *gin.Context) {
	start := time.Now()

	err := cache.DeleteCache("/data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	newData, exist := c.GetQuery("data")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "You must provide data form query.",
		})
		return
	}
	err = db.DataUpdate(newData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	duration := time.Since(start).Seconds()

	c.JSON(http.StatusAccepted, gin.H{
		"duration": duration,
	})
}
