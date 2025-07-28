package main

import (
	"log"
	"net/http"
	"os"

	"github.com/albqvictor1508/portfolio/cmd/db"
	"github.com/albqvictor1508/portfolio/internal"
	"github.com/albqvictor1508/portfolio/internal/project"
	"github.com/gin-gonic/gin"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	conn, err := db.NewConnection(databaseURL)
	if err != nil {
		log.Fatal("ERROR TO CONNECT TO DATABASE: ", err)
	}
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
		var project internal.Project

		if err := ctx.BindJSON(&project); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			return
		}

		project, err := repo.Insert(project)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":   true,
				"message": err.Error(),
			})
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"project": project,
		})
		return
	})

	if err := g.Run(":3333"); err != nil {
		log.Fatal("Fail to initialize server: ", err)
	}
}
