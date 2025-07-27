package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	_, err := db.NewConnection(databaseUrl)
	if err != nil {
		log.Fatal("ERROR TO CONNECT TO DATABASE: ", err)
	}
	g := gin.Default()

	g.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{ // gin.H é uma expressão que manda response, n sei pq
			"message": "Hello, World",
		})
	})

	if err := g.Run(":3333"); err != nil {
		log.Fatal("Fail to initialize server: ", err)
	}
}
