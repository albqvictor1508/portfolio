package function

import (
	"errors"
	"fmt"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/repository"
)

type TechnologyFunc struct {
	repo repository.TechnologyRepository
}

func NewTechnologyFunc(repo repository.TechnologyRepository) TechnologyFunc {
	return TechnologyFunc{
		repo: repo,
	}
}

func (cf *TechnologyFunc) CreateTechnology(technology *entity.Technology) (int, error) {
	existingTechnology, err := cf.repo.FindByName(technology.Name)
	if err != nil {
		return 0, fmt.Errorf("error checking for existing category: %w", err)
	}

	if existingTechnology.ID != 0 {
		return 0, errors.New("category with this name already exists")
	}

	return cf.repo.Insert(technology)
}

func (cf *TechnologyFunc) GetTechnologies() ([]entity.Technology, error) {
	return cf.repo.GetTechnologies()
}

func (cf *TechnologyFunc) GetTechnologyByID(id int) (entity.Technology, error) {
	return cf.repo.FindByID(id)
}

func (cf *TechnologyFunc) GetTechnologyByName(name string) (entity.Technology, error) {
	return cf.repo.FindByName(name)
}

func (cf *TechnologyFunc) DeleteByID(id int) error {
	return cf.repo.DeleteTechnologyByID(id)
}
