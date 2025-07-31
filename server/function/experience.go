package function

import (
	"errors"
	"fmt"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/repository"
)

type ExperienceFunction struct {
	experienceRepo repository.ExperienceRepository
	categoryRepo   repository.CategoryRepository
}

func NewExperienceFunc(experienceRepo repository.ExperienceRepository, categoryRepo repository.CategoryRepository) ExperienceFunction {
	return ExperienceFunction{
		experienceRepo: experienceRepo,
		categoryRepo:   categoryRepo,
	}
}

func (ef *ExperienceFunction) CreateExperience(e *entity.Experience) (int, error) {
	if len(e.CompanyName) < 3 {
		return 0, errors.New("the name must be at least 3 characters long")
	}

	if len(e.Description) < 100 {
		return 0, errors.New("the description must be at least 100 characters long")
	}

	if !isValidURL(e.PhotoURL) {
		return 0, errors.New("invalid github url")
	}

	experience, err := ef.experienceRepo.FindByName(e.CompanyName)
	if err != nil {
		return 0, fmt.Errorf("error finding project by name: %w", err)
	}

	if experience.ID != 0 {
		return 0, errors.New("a project with this name already exists")
	}

	if e.CategoryID != nil {
		category, err := ef.categoryRepo.FindByID(*e.CategoryID)
		if err != nil {
			return 0, fmt.Errorf("error finding category by id: %w", err)
		}
		if category.ID == 0 {
			return 0, errors.New("category not found")
		}
	}

	return ef.experienceRepo.Insert(e)
}

func (ef *ExperienceFunction) GetExperiences() ([]entity.Experience, error) {
	return ef.experienceRepo.GetExperiences()
}

func (pf *ExperienceFunction) DeleteExperience(id int) error {
	return pf.experienceRepo.Delete(id)
}

func (ef *ExperienceFunction) UpdateExperience(e *entity.Experience) (int, error) {
	experienceToUpdate, err := ef.experienceRepo.FindByID(e.ID)

	if experienceToUpdate.ID == 0 {
		return 0, errors.New("project not found")
	}

	if len(e.CompanyName) < 3 {
		return 0, errors.New("the name must be at least 3 characters long")
	}

	if len(e.Description) < 100 {
		return 0, errors.New("the description must be at least 100 characters long")
	}

	if !isValidURL(e.PhotoURL) {
		return 0, errors.New("invalid PhotoURL")
	}

	existingExperience, err := ef.experienceRepo.FindByName(e.CompanyName)
	if err != nil {
		return 0, fmt.Errorf("error finding experience by name: %w", err)
	}

	if existingExperience.ID != 0 && existingExperience.ID != p.ID {
		return 0, errors.New("a experience with this name already exists")
	}

	if e.CategoryID != nil {
		category, err := ef.categoryRepo.FindByID(*e.CategoryID)
		if err != nil {
			return 0, fmt.Errorf("error finding category by id: %w", err)
		}
		if category.ID == 0 {
			return 0, errors.New("category not found")
		}
	}

	return ef.experienceRepo.Update(e)
}
