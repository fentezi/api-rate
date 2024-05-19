package emails

import (
	"net/http"

	"github.com/fentezi/api-rate/internal/models"
	"github.com/fentezi/api-rate/internal/services"
	"github.com/gin-gonic/gin"
)

type EmailsController struct {
	emailsService *services.Service
}

func NewEmailsController(emailsService *services.Service) *EmailsController {
	return &EmailsController{emailsService: emailsService}
}

func (ec *EmailsController) Subscribe(c *gin.Context) {
	var email models.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := ec.emailsService.CreateEmail(email.Email)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if id != 0 {
		c.Status(http.StatusOK)
		return
	}
	c.Status(http.StatusInternalServerError)

}

func (ec *EmailsController) Mailing(c *gin.Context) {
	rate, err := ec.emailsService.Rate.GetRate()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = ec.emailsService.Mailing(rate)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
