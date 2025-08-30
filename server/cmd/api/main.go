package main

import (
	"log"
	"os"

	"github.com/albqvictor1508/portfolio/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	g := gin.Default()
	g.Use(cors.Default())

	g.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"health": "OK",
		})
	})
	g.POST("/contact", routes.SendEmail)
	g.POST("/upload", routes.UploadRoute)

	g.GET("/commits", routes.GetGithubData)
	g.GET("/cv", routes.SendCV)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	if err := g.Run(":" + port); err != nil {
		log.Fatal("Fail to initialize server: ", err)
	}
}
