package routes

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SendCV(ctx *gin.Context) {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("Error getting working directory:", err)
		ctx.AbortWithStatus(500)
		return
	}

	pdfPath := filepath.Join(wd, "public", "VICTOR_ALBUQUERQUE_CV.pdf")
	log.Println("Attempting to serve PDF from:", pdfPath)

	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", "attachment; filename=VICTOR_ALBUQUERQUE_CV.pdf")
	ctx.File(pdfPath)
}
