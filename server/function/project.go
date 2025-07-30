package function

import (
	"errors"
	"fmt"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/repository"
)

type ProjectFunction struct {
	projectRepo  repository.ProjectRepository
	categoryRepo repository.CategoryRepository
}

func NewProjectFunc(projectRepo repository.ProjectRepository, categoryRepo repository.CategoryRepository) ProjectFunction {
	return ProjectFunction{
		projectRepo:  projectRepo,
		categoryRepo: categoryRepo,
	}
}

func (pf *ProjectFunction) CreateProject(p *entity.Project) (int, error) {
	if len(p.Name) < 4 {
		return 0, errors.New("the name must be at least 4 characters long")
	}

	if len(p.Description) < 100 {
		return 0, errors.New("the description must be at least 100 characters long")
	}

	if !isValidURL(p.GithubURL) {
		return 0, errors.New("invalid github url")
	}

	if !isValidURL(p.DemoURL) {
		return 0, errors.New("invalid demo url")
	}

	project, err := pf.projectRepo.FindByName(p.Name)
	if err != nil {
		return 0, fmt.Errorf("error finding project by name: %w", err)
	}

	if project.ID != 0 {
		return 0, errors.New("a project with this name already exists")
	}

	if p.CategoryID != nil {
		category, err := pf.categoryRepo.FindByID(*p.CategoryID)
		if err != nil {
			return 0, fmt.Errorf("error finding category by id: %w", err)
		}
		if category.ID == 0 {
			return 0, errors.New("category not found")
		}
	}

	return pf.projectRepo.Insert(p)
}

func (pf *ProjectFunction) GetProjects() ([]entity.Project, error) {
	return pf.projectRepo.GetProjects()
}

func (pf *ProjectFunction) DeleteProject(id int) error {
	return pf.projectRepo.Delete(id)
}

func (pf *ProjectFunction) UpdateProject(p *entity.Project) (int, error) {
	projectToUpdate, err := pf.projectRepo.FindByID(p.ID)
	if err != nil {
		return 0, fmt.Errorf("error finding project by id: %w", err)
	}
	if projectToUpdate.ID == 0 {
		return 0, errors.New("project not found")
	}

	if len(p.Name) < 4 {
		return 0, errors.New("the name must be at least 4 characters long")
	}

	if len(p.Description) < 100 {
		return 0, errors.New("the description must be at least 100 characters long")
	}

	if !isValidURL(p.GithubURL) {
		return 0, errors.New("invalid github url")
	}

	if !isValidURL(p.DemoURL) {
		return 0, errors.New("invalid demo url")
	}

	existingProject, err := pf.projectRepo.FindByName(p.Name)
	if err != nil {
		return 0, fmt.Errorf("error finding project by name: %w", err)
	}

	if existingProject.ID != 0 && existingProject.ID != p.ID {
		return 0, errors.New("a project with this name already exists")
	}

	if p.CategoryID != nil {
		category, err := pf.categoryRepo.FindByID(*p.CategoryID)
		if err != nil {
			return 0, fmt.Errorf("error finding category by id: %w", err)
		}
		if category.ID == 0 {
			return 0, errors.New("category not found")
		}
	}

	return pf.projectRepo.Update(p)
}
