package controller

import "github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"

type DnaDto struct {
	Sequence [] string `json:"dna" binding:"required"`
}

func (dnaDto *DnaDto) ToModel() *domain.Dna {
	return domain.New(dnaDto.Sequence)
}