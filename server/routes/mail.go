package routes

import (
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
	err := utils.SendEmail(params.To)
}
