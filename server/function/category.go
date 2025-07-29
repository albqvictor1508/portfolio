package function

import (
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
	return cf.repo.Insert(category)
}

func (cf *CategoryFunc) GetProjects() ([]entity.Category, error) {
	return cf.repo.GetCategories()
}
