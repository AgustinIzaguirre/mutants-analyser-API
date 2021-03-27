package domain

import "github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"

type Service interface {
	GetStats() (Stats, errors.ApiError)
}
