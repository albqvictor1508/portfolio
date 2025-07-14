package repository

type CategoryRepository interface {

}

type categoryRepository struct {
	db *sql.DB
}

func (repository *CategoryRepository) CreateCategory(category *api.Category) (string, error) {
	query := `
	INSERT INTO categories (name) VALUES ($1)
	`

}
