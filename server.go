package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/echo", func(c *gin.Context) {
		param := c.Query("param")
		c.String(200, param)
	})

	r.GET("/read/:npm", readHandler)

	r.GET("/read/:npm/:trxId", readHandler)

	r.POST("/update", updateHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
