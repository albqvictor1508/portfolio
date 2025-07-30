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

	TechnologyRepository := repository.NewTechnology(conn)
	TechnologyFunction := function.NewTechnologyFunc(TechnologyRepository)
	TechnologyController := routes.NewTechnologyRoute(TechnologyFunction)

	g.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"health": "OK",
		})
	})
	g.GET("/projects", ProjectController.GetProjects)
	g.POST("/projects", ProjectController.CreateProject)
	g.PUT("/projects/:id", ProjectController.UpdateProject)
	g.DELETE("/projects/:id", ProjectController.DeleteProject)

	g.POST("/categories", CategoryController.CreateCategory)
	g.GET("/categories", CategoryController.GetCategories)
	g.GET("/categories/:id", CategoryController.FindByID)
	g.DELETE("/categories/:id", CategoryController.DeleteByID)

	g.POST("/technologies", TechnologyController.CreateTechnology)
	g.GET("/technologies", TechnologyController.GetTechnologies)
	g.GET("/technologies/:id", TechnologyController.FindByID)
	g.DELETE("/technologies/:id", TechnologyController.DeleteByID)
}
