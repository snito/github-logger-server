package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong11",
		})
	})
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong22311",
		})
	})	
	r.Run() // listen and serve on 0.0.0.0:8080
}