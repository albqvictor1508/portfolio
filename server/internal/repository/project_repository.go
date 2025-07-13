// Package repository contém a lógica de acesso ao banco de dados.
package repository

import (
	"database/sql"

	"github.com/albqvictor/server/cmd"
)

// ProjectRepository define a interface para as operações de banco de dados para projetos.
// Usar uma interface aqui é uma boa prática que permite "mockar" o repositório em testes.
type ProjectRepository interface {
	CreateProject(project *api.Project) (int64, error)
	GetProjectByID(id int64) (*api.Project, error)
	GetProjects() ([]*api.Project, error)
	UpdateProject(project *api.Project) error
	DeleteProject(id int64) error
}

// projectRepository é a implementação da interface ProjectRepository.
// Ela tem uma dependência do banco de dados (*sql.DB).
type projectRepository struct {
	db *sql.DB
}

// NewProjectRepository cria uma nova instância do repositório de projetos.
// Esta função é o que chamamos de "construtor".
func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db: db}
}

// CreateProject insere um novo projeto no banco de dados.
func (r *projectRepository) CreateProject(p *api.Project) (int64, error) {
	// A query SQL para inserir um novo projeto.
	// "RETURNING id" é usado para obter o ID do projeto recém-criado.
	query := `
		INSERT INTO projects (category_id, name, description, github_url, demo_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	var projectID int64
	// Executa a query, passando os dados do projeto como argumentos.
	// O método QueryRow é usado porque esperamos que a query retorne exatamente uma linha.
	err := r.db.QueryRow(query, p.CategoryID, p.Name, p.Description, p.GithubURL, p.DemoURL).Scan(&projectID)
	if err != nil {
		return 0, err
	}
	return projectID, nil
}

// GetProjectByID busca um projeto pelo seu ID.
func (r *projectRepository) GetProjectByID(id int64) (*api.Project, error) {
	// A query SQL para selecionar um projeto pelo ID.
	query := `
		SELECT id, category_id, name, description, github_url, demo_url, created_at, updated_at
		FROM projects
		WHERE id = $1
	`
	// Cria uma variável para armazenar o resultado.
	var p api.Project
	// Executa a query e faz o "Scan" dos resultados para dentro da struct 'p'.
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.CategoryID, &p.Name, &p.Description, &p.GithubURL, &p.DemoURL, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		// Se o erro for sql.ErrNoRows, significa que o projeto não foi encontrado.
		if err == sql.ErrNoRows {
			return nil, nil // Retorna nil para indicar que não foi encontrado.
		}
		return nil, err
	}
	return &p, nil
}

// GetProjects busca todos os projetos no banco de dados.
func (r *projectRepository) GetProjects() ([]*api.Project, error) {
	// A query SQL para selecionar todos os projetos.
	query := `
		SELECT id, category_id, name, description, github_url, demo_url, created_at, updated_at
		FROM projects
		ORDER BY created_at DESC
	`
	// Executa a query. O método Query é usado porque esperamos múltiplas linhas como resultado.
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Garante que as 'rows' sejam fechadas no final da função.
	defer rows.Close()

	// Cria um slice para armazenar os projetos.
	var projects []*api.Project
	// Itera sobre cada linha do resultado.
	for rows.Next() {
		var p api.Project
		// Faz o "Scan" dos dados da linha para a struct 'p'.
		if err := rows.Scan(&p.ID, &p.CategoryID, &p.Name, &p.Description, &p.GithubURL, &p.DemoURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		// Adiciona o projeto ao slice.
		projects = append(projects, &p)
	}

	return projects, nil
}

// UpdateProject atualiza os dados de um projeto existente.
func (r *projectRepository) UpdateProject(p *api.Project) error {
	// A query SQL para atualizar um projeto.
	query := `
		UPDATE projects
		SET category_id = $1, name = $2, description = $3, github_url = $4, demo_url = $5, updated_at = NOW()
		WHERE id = $6
	`
	// Executa a query. O método Exec é usado porque não esperamos que a query retorne linhas.
	result, err := r.db.Exec(query, p.CategoryID, p.Name, p.Description, p.GithubURL, p.DemoURL, p.ID)
	if err != nil {
		return err
	}

	// Verifica quantas linhas foram afetadas pela operação.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Se nenhuma linha foi afetada, significa que o projeto com o ID fornecido não existe.
	if rowsAffected == 0 {
		return sql.ErrNoRows // Usa um erro padrão para "não encontrado".
	}

	return nil
}

// DeleteProject remove um projeto do banco de dados pelo seu ID.
func (r *projectRepository) DeleteProject(id int64) error {
	// A query SQL para deletar um projeto.
	query := "DELETE FROM projects WHERE id = $1"
	// Executa a query.
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	// Verifica se alguma linha foi de fato deletada.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows // Retorna erro se o projeto não existia.
	}

	return nil
}
