// Package service contém a lógica de negócio da aplicação.
package service

import (
	"github.com/albqvictor/server/cmd"
	"github.com/albqvictor/server/internal/repository"
)

// ProjectService define a interface para a lógica de negócio relacionada a projetos.
// Assim como no repositório, usar uma interface aqui facilita os testes.
type ProjectService interface {
	CreateProject(project *api.Project) (int64, error)
	GetProjectByID(id int64) (*api.Project, error)
	GetProjects() ([]*api.Project, error)
	UpdateProject(project *api.Project) error
	DeleteProject(id int64) error
}

// projectService é a implementação da interface ProjectService.
// Ela depende do repositório de projetos.
type projectService struct {
	repo repository.ProjectRepository
}

// NewProjectService cria uma nova instância do serviço de projetos.
// Este é o "construtor" do nosso serviço.
func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

// CreateProject simplesmente repassa a chamada para o repositório.
// Em uma aplicação mais complexa, aqui você poderia adicionar validações,
// notificações, ou outra lógica de negócio antes de criar o projeto.
func (s *projectService) CreateProject(p *api.Project) (int64, error) {
	return s.repo.CreateProject(p)
}

// GetProjectByID busca um projeto pelo ID, repassando a chamada para o repositório.
func (s *projectService) GetProjectByID(id int64) (*api.Project, error) {
	return s.repo.GetProjectByID(id)
}

// GetProjects busca todos os projetos.
func (s *projectService) GetProjects() ([]*api.Project, error) {
	return s.repo.GetProjects()
}

// UpdateProject atualiza um projeto.
func (s *projectService) UpdateProject(p *api.Project) error {
	return s.repo.UpdateProject(p)
}

// DeleteProject deleta um projeto.
func (s *projectService) DeleteProject(id int64) error {
	return s.repo.DeleteProject(id)
}
