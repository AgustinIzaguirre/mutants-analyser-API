package domain

import "github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"

type Dao interface {
	GetStats() (Stats, errors.ApiError)
}
