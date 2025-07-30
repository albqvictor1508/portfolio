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

func (tf *TechnologyFunc) CreateTechnology(technology *entity.Technology) (int, error) {
	if !isValidURL(technology.PhotoURL) {
		return 0, errors.New("invalid photo url")
	}

	existingTechnology, err := tf.repo.FindByName(technology.Name)
	if err != nil {
		return 0, fmt.Errorf("error checking for existing technology: %w", err)
	}

	if existingTechnology.ID != 0 {
		return 0, errors.New("technology with this name already exists")
	}

	return tf.repo.Insert(technology)
}

func (tf *TechnologyFunc) GetTechnologies() ([]entity.Technology, error) {
	return tf.repo.GetTechnologies()
}

func (tf *TechnologyFunc) GetTechnologyByID(id int) (entity.Technology, error) {
	return tf.repo.FindByID(id)
}

func (tf *TechnologyFunc) GetTechnologyByName(name string) (entity.Technology, error) {
	return tf.repo.FindByName(name)
}

func (tf *TechnologyFunc) DeleteByID(id int) error {
	return tf.repo.DeleteTechnologyByID(id)
}
