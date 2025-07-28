package routes

import (
	"github.com/albqvictor1508/portfolio/function"
	"github.com/gin-gonic/gin"
)

type ProjectRoutes struct {
	function function.ProjectFunction
}

// dar uma olhada nessa questão
func (r *ProjectRoutes) ProjectRoutes() {
	g := gin.Default()
	g.POST("/projects", r.function.CreateProject)
}
