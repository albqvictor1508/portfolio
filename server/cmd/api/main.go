package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/albqvictor1508/portfolio/cmd/db"
	"github.com/albqvictor1508/portfolio/internal"
	"github.com/albqvictor1508/portfolio/internal/project"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	repo := project.RepositoryPg{
		Conn: conn,
	}

	g.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{ // gin.H é uma expressão que manda response, n sei pq
			"health": "OK",
		})
	})

	g.POST("/projects", func(ctx *gin.Context) {
		var project internal.InsertProjectParams

		if err := ctx.BindJSON(&project); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			return
		}

		newProject := internal.Project{
			ID:        uuid.New(),
			GithubURL: project.GithubURL,
			DemoURL:   project.DemoURL,
			IsPinned:  project.IsPinned,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		insertedProject, err := repo.Insert(newProject)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":   true,
				"message": err.Error(),
			})
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"project": insertedProject,
		})
	})

	if err := g.Run(":3333"); err != nil {
		log.Fatal("Fail to initialize server: ", err)
	}
}
