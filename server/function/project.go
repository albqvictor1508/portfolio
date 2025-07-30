package function

import (
	"errors"
	"fmt"
	"log"
	"net/http"

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
		return 0, errors.New("THE NAME MUST TO BE BETWEEN 3 CHARACTERS")
	}

	if len(p.Description) < 100 {
		return 0, errors.New("THE Description MUST TO BE BETWEEN 100 CHARACTERS")
	}

	if _, err := http.Get(p.GithubURL); p.GithubURL != "" && err != nil {
		errorMessage := fmt.Sprintf("INVALID GITHUB URL: %v", err)
		return 0, errors.New(errorMessage)
	}

	if _, err := http.Get(p.DemoURL); p.DemoURL != "" && err != nil {
		log.Fatal(err)
		return 0, errors.New("INVALID DEMO URL")
	}

	project, err := pf.projectRepo.FindByName(p.Name)
	if err != nil {
		errorMessage := fmt.Sprintf("ERROR TO FIND PROJECT BY NAME IN REPOSITORY: %v", err)
		return 0, errors.New(errorMessage)
	}

	if project != (entity.Project{}) {
		return 0, errors.New("PROJECT WITH THIS NAME ALREADY EXISTS")
	}
	if p.CategoryID != nil {
		category, err := pf.categoryRepo.FindByID(*p.CategoryID)
		if err != nil {
			errorMessage := fmt.Sprintf("ERROR TO FIND CATEGORY BY ID: %v", err)
			return 0, errors.New(errorMessage)
		}

		if category.ID == 0 {
			return 0, errors.New("CATEGORY WITH THIS ID NOT EXISTS")
		}
	}

	id, err := pf.projectRepo.Insert(p)
	if err != nil {
		errorMessage := fmt.Sprintf("ERROR TO INSERT PROJECT IN REPOSITORY: %v", err)
		return 0, errors.New(errorMessage)
	}
	return id, nil
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
		return 0, errors.New("THE NAME MUST TO BE BETWEEN 3 CHARACTERS")
	}

	if len(p.Description) < 100 {
		return 0, errors.New("THE Description MUST TO BE BETWEEN 100 CHARACTERS")
	}

	if p.GithubURL != "" {
		if _, err := http.Get(p.GithubURL); err != nil {
			return 0, fmt.Errorf("invalid github url: %w", err)
		}
	}

	if p.DemoURL != "" {
		if _, err := http.Get(p.DemoURL); err != nil {
			return 0, fmt.Errorf("invalid demo url: %w", err)
		}
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

	id, err := pf.projectRepo.Update(p)
	if err != nil {
		return 0, fmt.Errorf("error updating project: %w", err)
	}
	return id, nil
}
