package routes

import (
	"fmt"
	"net/http"

	"github.com/albqvictor1508/portfolio/utils"
	"github.com/gin-gonic/gin"
)

func SendEmail(ctx *gin.Context) {
	var params utils.SendEmailParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := utils.SendEmail(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("ERROR ON SEND EMAIL: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": params,
	})
}
