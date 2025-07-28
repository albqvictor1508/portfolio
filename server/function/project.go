package function

import (
	"fmt"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/repository"
)

type ProjectFunction struct {
	repo repository.Repository
}

func (pf *ProjectFunction) CreateProject(p *entity.Project) {
	fmt.Print("salve")
}
