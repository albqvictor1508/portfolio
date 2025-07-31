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

func (e *ExperienceRoutes) CreateExperience(ctx *gin.Context) {
	var experience *entity.Experience
	if err := ctx.ShouldBindJSON(&experience); err != nil {
		fmt.Errorf("Error to format request.body: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := e.experienceFunc.CreateExperience(experience)
	if err != nil {
		fmt.Print(err.Error())
		ctx.JSON(850, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (e *ExperienceRoutes) GetExperiences(ctx *gin.Context) {
	experiences, err := e.experienceFunc.GetExperiences()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"experiences": experiences,
	})
}

func (e *ExperienceRoutes) UpdateExperience(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid experience ID",
		})
		return
	}

	var experience entity.Experience
	if err := ctx.ShouldBindJSON(&experience); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	experience.ID = id

	if _, err := e.experienceFunc.UpdateExperience(&experience); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (e *ExperienceRoutes) DeleteExperience(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid experience ID",
		})
		return
	}

	if err := e.experienceFunc.DeleteExperience(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

