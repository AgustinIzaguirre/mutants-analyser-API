package service

import (
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"
)

type service struct {
	dao domain.Dao
}

func New(dao domain.Dao) domain.Service {
	return &service{dao: dao}
}

func (service *service) AddAnalysis(dna *domain.Dna, allowOverlapping bool) (bool, errors.ApiError) {
	if dna.IsValid() {
		analyser := domain.NewAnalyser(allowOverlapping)
		isMutant := analyser.IsMutant(dna.GetSequence())
		return service.dao.AddAnalysis(dna.ToString(), isMutant)
	} else {
		return false, errors.NewBadRequestError("Invalid DNA sequence")
	}
}
