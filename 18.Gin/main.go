package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running at 8080",
		})
	})

	r.GET("/about", func (c *gin.Context)  {
		c.JSON(200, gin.H) {
			"about": "this is me ayush sarodey"
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
