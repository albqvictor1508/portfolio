package function

import (
	"errors"
	"fmt"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/repository"
)

type CategoryFunc struct {
	repo repository.CategoryRepository
}

func NewCategoryFunc(repo repository.CategoryRepository) CategoryFunc {
	return CategoryFunc{
		repo: repo,
	}
}

func (cf *CategoryFunc) CreateCategory(category *entity.Category) (int, error) {
	existingCategory, err := cf.repo.FindByName(category.Name)
	if err != nil {
		return 0, fmt.Errorf("error checking for existing category: %w", err)
	}

	if existingCategory.ID != 0 {
		return 0, errors.New("category with this name already exists")
	}

	return cf.repo.Insert(category)
}

func (cf *CategoryFunc) GetCategories() ([]entity.Category, error) {
	return cf.repo.GetCategories()
}

func (cf *CategoryFunc) GetCategoryByID(id int) (entity.Category, error) {
	return cf.repo.FindByID(id)
}

func (cf *CategoryFunc) GetCategoryByName(name string) (entity.Category, error) {
	return cf.repo.FindByName(name)
}

/*
func (cf *CategoryFunc) GetCategoryByName(name string) error {
}
*/
