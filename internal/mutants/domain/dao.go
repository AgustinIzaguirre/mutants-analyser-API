package domain

type Dao interface {
	AddAnalysis(dna string, isMutant bool) (bool, error)
}
