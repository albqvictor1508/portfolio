package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/albqvictor1508/portfolio/common/images"
	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/function"
	"github.com/gin-gonic/gin"
)

type TechnologyRoutes struct {
	technologyFunc function.TechnologyFunc
}

func NewTechnologyRoute(technologyFunc function.TechnologyFunc) TechnologyRoutes {
	return TechnologyRoutes{
		technologyFunc: technologyFunc,
	}
}

func (tr *TechnologyRoutes) CreateTechnology(ctx *gin.Context) {
	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error":   err.Error(),
			"message": "'photo' is required, send this on form-data",
		})
	}
	var technology *entity.Technology
	technology.Name = ctx.PostForm("name")
	filename := strings.ReplaceAll(file.Filename, " ", "-")
	technologyName := strings.ReplaceAll(technology.Name, " ", "-")
	filepath := fmt.Sprintf("project/%v/%v", technologyName, filename)

	photoURL, err := images.UploadFile(file, filepath)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error":   err.Error(),
			"message": "Error to upload photo",
		})
	}
	technology.PhotoURL = photoURL
	if err := ctx.ShouldBindJSON(&technology); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := tr.technologyFunc.CreateTechnology(technology)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"technology_id": id,
	})
}

func (tr *TechnologyRoutes) GetTechnologies(ctx *gin.Context) {
	technologies, err := tr.technologyFunc.GetTechnologies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"technologies": technologies,
	})
}

func (tr *TechnologyRoutes) FindByID(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	technology, err := tr.technologyFunc.GetTechnologyByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if technology.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "technology with this id not exists",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"technology": technology,
	})
}

func (tr *TechnologyRoutes) DeleteByID(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := tr.technologyFunc.DeleteByID(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (tr *TechnologyRoutes) UpdateTechnology(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid technology ID",
		})
		return
	}

	var technology entity.Technology
	if err := ctx.ShouldBindJSON(&technology); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	technology.ID = id

	if _, err := tr.technologyFunc.UpdateTechnology(&technology); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
