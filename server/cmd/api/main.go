package main

import (
	"fmt"
	"log"
	"os"

	"github.com/albqvictor1508/portfolio/cmd/db"
	"github.com/albqvictor1508/portfolio/function"
	"github.com/albqvictor1508/portfolio/repository"
	"github.com/albqvictor1508/portfolio/routes"
	"github.com/gin-contrib/cors"
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
	g.Use(cors.Default())

	CategoryRepository := repository.NewCategory(conn)
	CategoryFunction := function.NewCategoryFunc(CategoryRepository)
	CategoryController := routes.NewCategoryRoute(CategoryFunction)

	ExperienceRepository := repository.NewExperience(conn)
	ExperienceFunc := function.NewExperienceFunc(ExperienceRepository, CategoryRepository)
	ExperienceController := routes.NewExperienceRoutes(ExperienceFunc)

	TechnologyRepository := repository.NewTechnology(conn)
	TechnologyFunction := function.NewTechnologyFunc(TechnologyRepository)
	TechnologyController := routes.NewTechnologyRoute(TechnologyFunction)

	ProjectRepository := repository.NewProjectRepo(conn)
	ProjectFunction := function.NewProjectFunc(ProjectRepository, CategoryRepository, TechnologyRepository)
	ProjectController := routes.NewProjectRoute(ProjectFunction)
	g.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"health": "OK",
		})
	})
	g.GET("/projects", ProjectController.GetProjects)
	g.POST("/projects", ProjectController.CreateProject)
	g.GET("/projects/:id", ProjectController.FindByID)
	g.PUT("/projects/:id", ProjectController.UpdateProject)
	g.DELETE("/projects/:id", ProjectController.DeleteProject)

	g.GET("/experiences", ExperienceController.GetExperiences)
	g.POST("/experiences", ExperienceController.CreateExperience)
	g.PUT("/experiences/:id", ExperienceController.UpdateExperience)
	g.DELETE("/experiences/:id", ExperienceController.DeleteExperience)

	g.POST("/categories", CategoryController.CreateCategory)
	g.GET("/categories", CategoryController.GetCategories)
	g.GET("/categories/:id", CategoryController.FindByID)
	g.DELETE("/categories/:id", CategoryController.DeleteByID)

	g.POST("/technologies", TechnologyController.CreateTechnology)
	g.GET("/technologies", TechnologyController.GetTechnologies)
	g.GET("/technologies/:id", TechnologyController.FindByID)
	g.DELETE("/technologies/:id", TechnologyController.DeleteByID)
	g.PUT("/technologies/:id", TechnologyController.UpdateTechnology)

	g.POST("/contact", routes.SendEmail)
	g.POST("/upload", routes.UploadRoute)

	g.GET("/commits", routes.GetGithubData)

	portStr := fmt.Sprintf(":%v", os.Getenv("PORT"))
	if err := g.Run(portStr); err != nil {
		log.Fatal("Fail to initialize server: ", err)
	}
}
