package rate

import (
	"net/http"

	"github.com/fentezi/api-rate/internal/services"
	"github.com/gin-gonic/gin"
)

type RateController struct {
	rateService *services.Service
}

func NewRateController(rateService *services.Service) *RateController {
	return &RateController{rateService: rateService}
}

func (rt *RateController) GetRate(c *gin.Context) {
	rate, err := rt.rateService.GetRate()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"rate": rate})
}
