package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {

}

func New() *Controller{
	return &Controller{}
}

func (controller *Controller) AnalyseDNA(context *gin.Context) {
	context.JSON(200, "Hello world")
}
