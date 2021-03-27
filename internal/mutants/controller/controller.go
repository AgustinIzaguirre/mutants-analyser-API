package controller

import (
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
)

type Controller struct {
	service domain.Service
}

func New(mutantService domain.Service) *Controller{
	return &Controller{service: mutantService}
}

func (controller *Controller) AnalyseDNA(context *gin.Context) {
	params := context.Request.URL.Query()
	allowOverllapping := getBoolFromQueryParam(params, "overlapping", true)
	var dna DnaDto
	if err := context.ShouldBind(&dna); err != nil {
		context.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
	} else {
		isMutant, err := controller.service.AddAnalysis(dna.ToModel(), allowOverllapping)
		if err != nil {
			context.JSON(err.GetStatus(), err)
		} else if !isMutant {
			context.JSON(http.StatusForbidden, errors.NewForbiddenError("Is Human"))
		} else {
			context.JSON(http.StatusOK, "Is Mutant")
		}
	}
}

func getBoolFromQueryParam(queryParams map[string][]string, paramName string, defaultValue bool) bool {
	valueArray := queryParams[paramName]
	if len(valueArray) == 0 {
		return defaultValue
	} else if strings.ToLower(valueArray[0]) == "false" {
		return false
	} else if strings.ToLower(valueArray[0]) == "true" {
		return true
	} else {
		return defaultValue
	}
}

