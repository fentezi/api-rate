package controllers

import (
	"github.com/fentezi/api-rate/internal/controllers/emails"
	"github.com/fentezi/api-rate/internal/controllers/rate"
	"github.com/fentezi/api-rate/internal/services"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Emails
	Rate
}

type Emails interface {
	Subscribe(c *gin.Context)
	Mailing(c *gin.Context)
}

type Rate interface {
	GetRate(c *gin.Context)
}

func NewController(service *services.Service) *Controller {
	return &Controller{
		Emails: emails.NewEmailsController(service),
		Rate:   rate.NewRateController(service),
	}
}
