package api

import (
	"go-redis-example/internal/app/api/cache"
	"go-redis-example/internal/app/api/router"
)

// Run start api server
func Run() {
	cache.InitRedis("localhost:6379", "", 0)

	web := router.Setup()
	panic(web.Run(":10912"))
}
