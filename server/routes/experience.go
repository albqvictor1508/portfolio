package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/albqvictor1508/portfolio/common/images"
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
	var experience entity.Experience

	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "'photo' is required, please send them on form-data",
		})
	}

	categoryIDStr := ctx.PostForm("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "category_id is required",
		})
	}
	experience.CompanyName = ctx.PostForm("company_name")
	experience.CategoryID = &categoryID
	experience.Role = ctx.PostForm("role")
	experience.Description = ctx.PostForm("description")

	startDateStr := ctx.PostForm("start_date")
	endDateStr := ctx.PostForm("end_date")

	experience.StartDate, err = time.Parse("YYYY-MM-DD", startDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid date",
		})
		experience.EndDate, err = time.Parse("YYYY-MM-DD", endDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "invalid date",
			})
		}
	}
	filename := strings.ReplaceAll(file.Filename, " ", "-")
	experienceName := strings.ReplaceAll(experience.CompanyName, " ", "-")
	filePath := fmt.Sprintf("project/%v/%v", experienceName, filename)

	photoURL, err := images.UploadFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "error to upload file",
		})
	}

	experience.PhotoURL = &photoURL

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

	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "'photo' is required, please send them on form-data",
		})
	}

	categoryIDStr := ctx.PostForm("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "category_id is required",
		})
	}
	experience.CompanyName = ctx.PostForm("company_name")
	experience.CategoryID = &categoryID
	experience.Role = ctx.PostForm("role")
	experience.Description = ctx.PostForm("description")

	startDateStr := ctx.PostForm("start_date")
	endDateStr := ctx.PostForm("end_date")

	experience.StartDate, err = time.Parse("YYYY-MM-DD", startDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid date",
		})
		experience.EndDate, err = time.Parse("YYYY-MM-DD", endDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "invalid date",
			})
		}
	}
	filename := strings.ReplaceAll(file.Filename, " ", "-")
	experienceName := strings.ReplaceAll(experience.CompanyName, " ", "-")
	filePath := fmt.Sprintf("project/%v/%v", experienceName, filename)

	photoURL, err := images.UploadFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "error to upload file",
		})
	}

	experience.PhotoURL = &photoURL
	experience.ID = id

	if _, err := e.experienceFunc.UpdateExperience(&experience); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error to update experience"})
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
