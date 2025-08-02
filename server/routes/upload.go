package routes

import (
	"net/http"

	"github.com/albqvictor1508/portfolio/common/images"
	"github.com/gin-gonic/gin"
)

func UploadRoute(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Missing file, send them on form-data with name 'file'",
		})
	}
	if _, err := images.UploadFile(file); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error to upload file on R2 Cloudflare Bucket",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
