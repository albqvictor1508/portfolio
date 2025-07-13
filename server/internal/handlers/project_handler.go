// Package handlers contém os handlers HTTP da nossa API.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/albqvictor1508/server/cmd"
	"github.com/albqvictor1508/server/internal/service"
	"github.com/go-chi/chi/v5"
)

// ProjectHandler é a estrutura que lida com as requisições HTTP para projetos.
// Ele depende do serviço de projetos.
type ProjectHandler struct {
	service service.ProjectService
}

// NewProjectHandler cria um novo handler de projetos.
func NewProjectHandler(service service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

// CreateProject é o handler para a rota POST /projects.
// Ele decodifica o JSON do corpo da requisição e cria um novo projeto.
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var p api.Project
	// Decodifica o JSON do corpo da requisição para a struct Project.
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		// Se houver um erro na decodificação, retorna um erro 400 (Bad Request).
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Chama o serviço para criar o projeto.
	projectID, err := h.service.CreateProject(&p)
	if err != nil {
		// Se o serviço retornar um erro, retorna um erro 500 (Internal Server Error).
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	// Define o ID do projeto criado para a resposta.
	p.ID = projectID

	// Define o cabeçalho da resposta como JSON.
	w.Header().Set("Content-Type", "application/json")
	// Define o status da resposta como 201 (Created).
	w.WriteHeader(http.StatusCreated)
	// Codifica o projeto criado como JSON e o envia na resposta.
	json.NewEncoder(w).Encode(p)
}

// GetProjectByID é o handler para a rota GET /projects/{id}.
// Ele extrai o ID da URL e busca o projeto correspondente.
func (h *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o parâmetro "id" da URL usando o chi router.
	idStr := chi.URLParam(r, "id")
	// Converte o ID de string para int64.
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Se o ID não for um número válido, retorna um erro 400.
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Chama o serviço para buscar o projeto.
	project, err := h.service.GetProjectByID(id)
	if err != nil {
		// Se o serviço retornar um erro, retorna um erro 500.
		http.Error(w, "Failed to get project", http.StatusInternalServerError)
		return
	}

	// Se o projeto não for encontrado, retorna um erro 404 (Not Found).
	if project == nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Responde com o projeto encontrado em formato JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

// GetProjects é o handler para a rota GET /projects.
// Ele busca todos os projetos.
func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	// Chama o serviço para buscar todos os projetos.
	projects, err := h.service.GetProjects()
	if err != nil {
		// Se houver um erro, retorna um erro 500.
		http.Error(w, "Failed to get projects", http.StatusInternalServerError)
		return
	}

	// Responde com a lista de projetos em formato JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// UpdateProject é o handler para a rota PUT /projects/{id}.
// Ele atualiza um projeto existente.
func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	var p api.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Define o ID do projeto a ser atualizado.
	p.ID = id

	// Chama o serviço para atualizar o projeto.
	if err := h.service.UpdateProject(&p); err != nil {
		// Aqui, poderíamos verificar o tipo de erro para retornar 404 se não encontrado.
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta 204 (No Content) para indicar sucesso sem corpo de resposta.
	w.WriteHeader(http.StatusNoContent)
}

// DeleteProject é o handler para a rota DELETE /projects/{id}.
// Ele deleta um projeto.
func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Chama o serviço para deletar o projeto.
	if err := h.service.DeleteProject(id); err != nil {
		// Poderíamos verificar o erro para retornar 404 se não encontrado.
		http.Error(w, "Failed to delete project", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta 204 (No Content) para indicar sucesso.
	w.WriteHeader(http.StatusNoContent)
}
