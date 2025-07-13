package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/albqvictor/server/internal/handlers"
	"github.com/albqvictor/server/internal/repository"
	"github.com/albqvictor/server/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3" // O _ anônimo importa o driver do SQLite para que ele se registre com o pacote database/sql.
)

func main() {
	// --- 1. Conexão com o Banco de Dados ---
	// Usamos o SQLite para simplicidade. Ele cria um arquivo no disco.
	// O `?_foreign_keys=on` é importante para garantir a integridade dos dados.
	db, err := sql.Open("sqlite3", "./database.db?_foreign_keys=on")
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	// Garante que a conexão com o banco seja fechada ao final da função main.
	defer db.Close()

	// --- 2. Criação da Tabela (para fins de exemplo) ---
	// Este trecho cria a tabela `projects` se ela ainda não existir.
	// Em um projeto real, você usaria um sistema de migrations como o que já tem na pasta /migrations.
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

	// --- 3. Injeção de Dependências ---
	// Aqui, "injetamos" as dependências de baixo para cima.
	// O handler precisa do serviço, e o serviço precisa do repositório.
	// O repositório precisa da conexão com o banco de dados.
	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)

	// --- 4. Configuração do Roteador (chi) ---
	// Cria uma nova instância do roteador chi.
	r := chi.NewRouter()

	// Adiciona middlewares. Middlewares são executados em cada requisição.
	r.Use(middleware.Logger) // Loga informações sobre cada requisição (método, path, tempo de resposta).
	r.Use(middleware.Recoverer) // Recupera de panics para que o servidor não quebre.

	// --- 5. Definição das Rotas de Projetos ---
	// Agrupa as rotas que começam com /projects.
	r.Route("/projects", func(r chi.Router) {
		r.Post("/", projectHandler.CreateProject)       // POST /projects
		r.Get("/", projectHandler.GetProjects)         // GET /projects
		r.Get("/", projectHandler.GetProjects)         // GET /projects
		r.Get("/{id}", projectHandler.GetProjectByID) // GET /projects/123
		r.Put("/{id}", projectHandler.UpdateProject)    // PUT /projects/123
		r.Delete("/{id}", projectHandler.DeleteProject) // DELETE /projects/123
	})

	// --- 6. Início do Servidor ---
	log.Println("Servidor iniciando na porta :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}