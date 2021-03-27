package controller

import (
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/stats/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service domain.Service
}

func New(mutantService domain.Service) *Controller{
	return &Controller{service: mutantService}
}

func (controller *Controller) GetStats(context *gin.Context) {
	stats, err := controller.service.GetStats()
	if err != nil {
		context.JSON(500, err)
	}
	context.JSON(http.StatusOK, stats)
}
