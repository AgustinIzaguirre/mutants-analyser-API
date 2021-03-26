package service

import "github.com/AgustinIzaguirre/mutants-analyser-api/internal/stats/domain"


type service struct {
	dao domain.Dao
}

func New(dao domain.Dao) domain.Service {
	return &service{dao: dao}
}

func (service *service) GetStats() (domain.Stats, error) {
	return service.dao.GetStats()
}

