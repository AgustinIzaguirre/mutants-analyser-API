package domain

type Service interface {
	AddAnalysis(dna *Dna, allowOverlapping bool) (bool, error)
}
