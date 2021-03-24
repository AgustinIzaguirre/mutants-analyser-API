package controller

import (
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	dao domain.Dao
}

func New(mutantDao domain.Dao) *Controller{
	return &Controller{dao: mutantDao}
}

func (controller *Controller) AnalyseDNA(context *gin.Context) {
	_ = controller.dao.AddAnalysis(true)
	context.JSON(200, "Hello world")
}
