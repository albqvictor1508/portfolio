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
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("postgres", "postgres://albqvxc:lexsa1508@localhost:6005/portfolio?sslmode=disable")
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	m, err := migrate.New(
		"file://migrations",
		"postgres://albqvxc:lexsa1508@localhost:6005/portfolio?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
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
