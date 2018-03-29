package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
)

func main() {
	count := 1
	callUuidV4, _ := uuid.NewV4()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		instanceUuidV4, _ := uuid.NewV4()

		c.JSON(200, gin.H{
			"message":       "ok",
			"time":          time.Now(),
			"count":         count,
			"uuid_call":     instanceUuidV4.String(),
			"uuid_instance": callUuidV4.String(),
			"client_ip":     c.ClientIP(),
			"version":       2,
			"version_msg":   "version 2",
		})
		count++
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
