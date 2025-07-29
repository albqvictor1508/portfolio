package function

import "github.com/albqvictor1508/portfolio/repository"

type CategoryFunc struct {
	repo repository.CategoryRepository
}

func NewCategoryFunc(repo *repository.CategoryRepository) CategoryFunc {
	return CategoryFunc{
		repo: *repo,
	}
}

func (cf *CategoryFunc) CreateCategory() {
}
