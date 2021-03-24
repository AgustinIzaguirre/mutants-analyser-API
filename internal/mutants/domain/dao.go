package domain

type Dao interface {
	AddAnalysis(isMutant bool) error
}
