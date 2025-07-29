package main

import (
	"log"
	"os"

	"github.com/albqvictor1508/portfolio/cmd/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Godotenv not loaded")
	}
	databaseURL := os.Getenv("DATABASE_URL")

	conn, err := db.NewConnection(databaseURL)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	g := gin.Default()

	g.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{ // gin.H é uma expressão que manda response, n sei pq
			"health": "OK",
		})
	})

	if err := g.Run(":3333"); err != nil {
		log.Fatal("Fail to initialize server: ", err)
	}
}
