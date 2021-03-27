package domain

import "github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"

type Service interface {
	AddAnalysis(dna *Dna, allowOverlapping bool) (bool, errors.ApiError)
}
