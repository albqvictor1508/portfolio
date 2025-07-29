package main

import (
	"log"
	"os"

	"github.com/albqvictor1508/portfolio/cmd/db"
	"github.com/albqvictor1508/portfolio/function"
	"github.com/albqvictor1508/portfolio/repository"
	"github.com/albqvictor1508/portfolio/routes"
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

	ProjectRepository := repository.New(conn)
	ProjectFunction := function.New(ProjectRepository)
	ProjectController := routes.New(ProjectFunction)

	g.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"health": "OK",
		})
	})
	g.GET("/projects", ProjectController.GetProjects)
	g.POST("/projects", ProjectController.CreateProject)

	if err := g.Run(":3333"); err != nil {
		log.Fatal("Fail to initialize server: ", err)
	}
}
