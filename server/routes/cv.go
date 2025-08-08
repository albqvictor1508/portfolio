package routes

import "github.com/gin-gonic/gin"

func SendCV(ctx *gin.Context) {
	ctx.Header("Content-Disposition", "attachment; filename=VICTOR_ALBUQUERQUE_CV.pdf")
	ctx.File("../public/cv.pdf")
}
