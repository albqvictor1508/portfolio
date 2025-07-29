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
	projectRepo repository.ProjectRepository
}

func New(repo repository.ProjectRepository) ProjectFunction {
	return ProjectFunction{
		projectRepo: repo,
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

	// validar se o projeto já existe pelo nome q é unique
	_, err := pf.projectRepo.FindByName(p.Name)
	if err != nil {
		errorMessage := fmt.Sprintf("ERROR TO FIND PROJECT BY NAME IN REPOSITORY: %v", err)
		return 0, errors.New(errorMessage)
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
