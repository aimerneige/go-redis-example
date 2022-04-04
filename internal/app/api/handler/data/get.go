package data

import (
	"go-redis-example/internal/app/api/cache"
	"go-redis-example/internal/app/api/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Get handler for get data
// Method: GET
// Path:   `/data/`
func Get(c *gin.Context) {
	start := time.Now()

	data, err := cache.GetKeyOrSetCache("/data", queryCacheFromDatabase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	duration := time.Since(start).Seconds()

	c.JSON(http.StatusOK, gin.H{
		"data":     data,
		"duration": duration,
	})
}

func queryCacheFromDatabase() (string, error) {
	return db.SlowQueryGet()
}
