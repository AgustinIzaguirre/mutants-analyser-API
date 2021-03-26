package service

import "github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"

type service struct {
	dao domain.Dao
}

func New(dao domain.Dao) domain.Service {
	return &service{dao: dao}
}

func (service *service) AddAnalysis(dna *domain.Dna, allowOverlapping bool) (bool, error) {
	if dna.IsValid() {
		analyser := domain.NewAnalyser(allowOverlapping)
		isMutant := analyser.IsMutant(dna.GetSequence())
		return service.dao.AddAnalysis(dna.ToString(), isMutant)
	} else {
		// TODO implement invalid DNA sequence or maybe BAD Request
		return false, nil
	}
}
