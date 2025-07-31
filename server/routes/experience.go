package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/function"
	"github.com/gin-gonic/gin"
)

type ExperienceRoutes struct {
	experienceFunc function.ExperienceFunction
}

func NewExperienceRoutes(experienceFunc function.ExperienceFunction) ExperienceRoutes {
	return ExperienceRoutes{
		experienceFunc: experienceFunc,
	}
}

func (p *projectRoute) CreateProject(ctx *gin.Context) {
	var project *entity.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := p.projectFunc.CreateProject(project)
	if err != nil {
		fmt.Print(err.Error())
		ctx.JSON(850, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (p *projectRoute) GetProjects(ctx *gin.Context) {
	projects, err := p.projectFunc.GetProjects()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func (p *projectRoute) UpdateProject(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid project ID",
		})
		return
	}

	var project entity.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	project.ID = id

	if _, err := p.projectFunc.UpdateProject(&project); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (p *projectRoute) DeleteProject(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid project ID",
		})
		return
	}

	if err := p.projectFunc.DeleteProject(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
