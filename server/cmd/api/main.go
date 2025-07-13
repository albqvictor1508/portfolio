package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/albqvictor1508/server/internal/handlers"
	"github.com/albqvictor1508/server/internal/repository"
	"github.com/albqvictor1508/server/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-postgres"
)

func main() {
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

createTableSQL := `
	CREATE TABLE IF NOT EXISTS projects (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"category_id" INTEGER,
		"name" TEXT,
		"description" TEXT,
		"github_url" TEXT,
		"demo_url" TEXT,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
		"updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela: %v", err)
	}

	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)

	r := chi.NewRouter()

	r.Use(middleware.Logger) 
	r.Use(middleware.Recoverer)

	r.Route("/projects", func(r chi.Router) {
		r.Post("/", projectHandler.CreateProject)       
		r.Get("/", projectHandler.GetProjects)         
		r.Get("/", projectHandler.GetProjects)         
		r.Get("/{id}", projectHandler.GetProjectByID) 
		r.Put("/{id}", projectHandler.UpdateProject)    
		r.Delete("/{id}", projectHandler.DeleteProject) 
	})

	log.Println("Servidor iniciando na porta :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
