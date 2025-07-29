package routes

import (
	"fmt"
	"net/http"

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

func (cr *CategoryRoutes) CreateCategory(ctx *gin.Context, category *entity.Category) {
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

func (cr *CategoryRoutes) FindByID(ctx *gin.Context) {
}
