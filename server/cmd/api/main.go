package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{ // gin.H é uma expressão que manda response, n sei pq
			"message": "Hello, World",
		})
	})

	g.Run(":3333")
}
