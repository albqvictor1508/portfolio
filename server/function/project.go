package function

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/repository"
)

type ProjectFunction struct {
	projectRepo repository.Repository
}

func (pf *ProjectFunction) CreateProject(p *entity.Project) (int, error) {
	if len(p.Name) < 4 {
		return 0, errors.New("THE NAME MUST TO BE BETWEEN 3 CHARACTERS")
	}

	if len(p.Description) < 100 {
		return 0, errors.New("THE Description MUST TO BE BETWEEN 100 CHARACTERS")
	}

	if _, err := http.Get(p.GithubURL); err != nil {
		fmt.Println(err)
		return 0, errors.New("INVALID GITHUB URL")
	}

	if _, err := http.Get(p.GithubURL); err != nil {
		fmt.Println(err)
		return 0, errors.New("INVALID DEMO URL")
	}

	// validar se o projeto já existe pelo nome q é unique
	pf.projectRepo.
		pf.projectRepo.Insert(p)
	fmt.Print("salve")
}
