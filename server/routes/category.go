package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/albqvictor1508/portfolio/function"
	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	categoryFunc function.CategoryFunc
}

func NewCategoryRoute(categoryFunc function.CategoryFunc) CategoryRoutes {
	return CategoryRoutes{
		categoryFunc: categoryFunc,
	}
}

func (cr *CategoryRoutes) CreateCategory(ctx *gin.Context) {
	var category *entity.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := cr.categoryFunc.CreateCategory(category)
	if err != nil {
		errorMessage := fmt.Sprintf("ERROR ON CREATE CATEGORY: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"category_id": id,
	})
}

func (cr *CategoryRoutes) GetCategories(ctx *gin.Context) {
	categories, err := cr.categoryFunc.GetCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (cr *CategoryRoutes) FindByID(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	category, err := cr.categoryFunc.GetCategoryByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if category.ID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "category with this id not exists",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func (cr *CategoryRoutes) DeleteByID(ctx *gin.Context) {
	vars := ctx.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := cr.categoryFunc.DeleteByID(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
