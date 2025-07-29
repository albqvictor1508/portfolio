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

	CategoryRepository := repository.NewCategory(conn)
	CategoryFunction := function.NewCategoryFunc(CategoryRepository)
	CategoryController := routes.NewCategoryRoute(CategoryFunction)

	ProjectRepository := repository.NewProjectRepo(conn)
	ProjectFunction := function.NewProjectFunc(ProjectRepository, CategoryRepository)
	ProjectController := routes.NewProjectRoute(ProjectFunction)

	g.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"health": "OK",
		})
	})
	g.GET("/projects", ProjectController.GetProjects)
	g.POST("/projects", ProjectController.CreateProject)

	g.POST("/categories", CategoryController.CreateCategory)
	g.GET("/categories", CategoryController.GetCategories)

	if err := g.Run(":3333"); err != nil {
		log.Fatal("Fail to initialize server: ", err)
	}
}
