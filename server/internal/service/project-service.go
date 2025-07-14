package service

import (
	"github.com/albqvictor1508/server/cmd"
	"github.com/albqvictor1508/server/internal/repository"
)

type ProjectService interface {
	CreateProject(project *api.Project) (string, error)
	GetProjectByID(id string) (*api.Project, error)
	GetProjects() ([]*api.Project, error)
	UpdateProject(project *api.Project) error
	DeleteProject(id string) error
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) CreateProject(p *api.Project) (string, error) {
	return s.repo.CreateProject(p)
}

func (s *projectService) GetProjectByID(id string) (*api.Project, error) {
	return s.repo.GetProjectByID(id)
}

func (s *projectService) GetProjects() ([]*api.Project, error) {
	return s.repo.GetProjects()
}

func (s *projectService) UpdateProject(p *api.Project) error {
	return s.repo.UpdateProject(p)
}

func (s *projectService) DeleteProject(id string) error {
	return s.repo.DeleteProject(id)
}
