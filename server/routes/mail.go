package routes

import (
	"fmt"
	"net/http"

	"github.com/albqvictor1508/portfolio/utils"
	"github.com/gin-gonic/gin"
)

type MailRoutes struct {
	params utils.SendEmailParams
}

func NewMailRoutes(params utils.SendEmailParams) MailRoutes {
	return MailRoutes{
		params: params,
	}
}

func (mr *MailRoutes) SendEmail(ctx *gin.Context) {
	var params *utils.SendEmailParams
	err := utils.SendEmail(*params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("ERROR TO SEND EMAIL: %v", err),
		}),
		return
	}
}
