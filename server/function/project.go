package function

import (
	"errors"
	"fmt"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/repository"
	"github.com/albqvictor1508/portfolio/utils"
)

type ProjectFunction struct {
	projectRepo    repository.ProjectRepository
	categoryRepo   repository.CategoryRepository
	technologyRepo repository.TechnologyRepository
}

func NewProjectFunc(projectRepo repository.ProjectRepository, categoryRepo repository.CategoryRepository, technologyRepo repository.TechnologyRepository) ProjectFunction {
	return ProjectFunction{
		projectRepo:    projectRepo,
		categoryRepo:   categoryRepo,
		technologyRepo: technologyRepo,
	}
}

func (pf *ProjectFunction) CreateProject(p *entity.Project) (*entity.Project, error) {
	if len(p.Name) < 4 {
		return nil, errors.New("the name must be at least 4 characters long")
	}

	if len(p.Description) < 100 {
		return nil, errors.New("the description must be at least 100 characters long")
	}

	if !utils.IsValidURL(p.GithubURL) {
		return nil, errors.New("invalid github url")
	}

	if p.DemoURL != "" && !utils.IsValidURL(p.DemoURL) {
		return nil, errors.New("invalid demo url")
	}

	project, err := pf.projectRepo.FindByName(p.Name)
	if err != nil {
		return nil, fmt.Errorf("error finding project by name: %w", err)
	}

	if project.ID != 0 {
		return nil, errors.New("a project with this name already exists")
	}

	if p.Category != nil && p.Category.ID != 0 {
		category, err := pf.categoryRepo.FindByID(p.Category.ID)
		if err != nil {
			return nil, fmt.Errorf("error finding category by id: %w", err)
		}
		if category.ID == 0 {
			return nil, errors.New("category not found")
		}
	}

	if len(p.Technologies) > 0 {
		for _, tech := range p.Technologies {
			technology, err := pf.technologyRepo.FindByID(tech.ID)
			if err != nil {
				return nil, fmt.Errorf("error finding technology by id: %w", err)
			}
			if technology.ID == 0 {
				return nil, fmt.Errorf("technology with id %d not found", tech.ID)
			}
		}
	}

	newID, err := pf.projectRepo.Insert(p)
	if err != nil {
		return nil, err
	}

	newProject, err := pf.projectRepo.FindByID(newID)
	if err != nil {
		return &entity.Project{}, fmt.Errorf("error to find by id: %v", err)
	}
	return &newProject, nil
}

func (pf *ProjectFunction) GetProjects() ([]entity.Project, error) {
	return pf.projectRepo.GetProjects()
}

func (pf *ProjectFunction) DeleteProject(id int) error {
	return pf.projectRepo.Delete(id)
}

func (pf *ProjectFunction) UpdateProject(p *entity.Project) (*entity.Project, error) {
	projectToUpdate, err := pf.projectRepo.FindByID(p.ID)
	if err != nil {
		return nil, fmt.Errorf("error finding project by id: %w", err)
	}
	if projectToUpdate.ID == 0 {
		return nil, errors.New("project not found")
	}

	if len(p.Name) < 4 {
		return nil, errors.New("the name must be at least 4 characters long")
	}

	if len(p.Description) < 100 {
		return nil, errors.New("the description must be at least 100 characters long")
	}

	if !utils.IsValidURL(p.GithubURL) {
		return nil, errors.New("invalid github url")
	}

	if p.DemoURL != "" && !utils.IsValidURL(p.DemoURL) {
		return nil, errors.New("invalid demo url")
	}

	existingProject, err := pf.projectRepo.FindByName(p.Name)
	if err != nil {
		return nil, fmt.Errorf("error finding project by name: %w", err)
	}

	if existingProject.ID != 0 && existingProject.ID != p.ID {
		return nil, errors.New("a project with this name already exists")
	}

	if p.Category != nil && p.Category.ID != 0 {
		category, err := pf.categoryRepo.FindByID(p.Category.ID)
		if err != nil {
			return nil, fmt.Errorf("error finding category by id: %w", err)
		}
		if category.ID == 0 {
			return nil, errors.New("category not found")
		}
	}

	if len(p.Technologies) > 0 {
		for _, tech := range p.Technologies {
			technology, err := pf.technologyRepo.FindByID(tech.ID)
			if err != nil {
				return nil, fmt.Errorf("error finding technology by id: %w", err)
			}
			if technology.ID == 0 {
				return nil, fmt.Errorf("technology with id %d not found", tech.ID)
			}
		}
	}

	updatedID, err := pf.projectRepo.Update(p)
	if err != nil {
		return nil, err
	}

	updatedProject, err := pf.projectRepo.FindByID(updatedID)
	if err != nil {
		return &entity.Project{}, fmt.Errorf("error to find by id: %v", err)
	}

	return &updatedProject, nil
}

func (pf *ProjectFunction) FindByID(id int) (entity.Project, error) {
	return pf.projectRepo.FindByID(id)
}
