package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	router.GET("hello/:name", func(c *gin.Context) {
		// Path Parameters
		name := c.Param("name")

		// Query Strings
		loud := c.DefaultQuery("loud", "false")

		greeting := "hello, " + name

		if loud == "true" {
			greeting = strings.ToUpper(greeting)
		}

		c.JSON(200, gin.H{"greeting": greeting})
	})

	err := router.Run(":8080")
	if err != nil {
		println("The router encountered an error\n", err)
	}
}
