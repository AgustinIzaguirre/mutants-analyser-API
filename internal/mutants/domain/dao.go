package domain

import "github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"

type Dao interface {
	AddAnalysis(dna string, isMutant bool) (bool, errors.ApiError)
}
