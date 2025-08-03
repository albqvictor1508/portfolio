package routes

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/albqvictor1508/portfolio/common/images"
	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/function"
	"github.com/gin-gonic/gin"
)

type projectRoute struct {
	projectFunc function.ProjectFunction
}

func NewProjectRoute(projectFunc function.ProjectFunction) projectRoute {
	return projectRoute{
		projectFunc: projectFunc,
	}
}

func (p *projectRoute) CreateProject(ctx *gin.Context) {
	// 1. Handle the file upload first
	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "'photo' is required, send it as form-data with the name 'photo'",
		})
		return
	}

	photoURL, uploadErr := images.UploadFile(file)
	if uploadErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   uploadErr.Error(),
			"message": "Error uploading photo to Cloudflare R2 Bucket",
		})
		return
	}

	var project entity.Project
	project.Name = ctx.PostForm("name")
	project.Description = ctx.PostForm("description")
	project.GithubURL = ctx.PostForm("github_url")
	project.DemoURL = ctx.PostForm("demo_url")
	project.PhotoURL = photoURL

	categoryIDStr := ctx.PostForm("category_id")
	if categoryIDStr != "" {
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'category_id' format, must be an integer"})
			return
		}
		project.Category = &entity.Category{ID: categoryID} // Correctly assign the category object
	}

	isPinned, _ := strconv.ParseBool(ctx.PostForm("is_pinned"))
	project.IsPinned = isPinned

	techIDsStr := ctx.PostForm("technologies")
	if techIDsStr != "" {
		for _, idStr := range strings.Split(techIDsStr, ",") { // Corrected to strings.Split
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'technologies' format, must be a comma-separated string of integers"})
				return
			}
			project.Technologies = append(project.Technologies, entity.Technology{ID: id})
		}
	}

	// The function layer now returns the created project object
	createdProject, err := p.projectFunc.CreateProject(&project)
	if err != nil {
		// Assuming the function layer provides specific error messages
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdProject)
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

func (p *projectRoute) FindByID(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid project ID",
		})
		return
	}

	project, err := p.projectFunc.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}
