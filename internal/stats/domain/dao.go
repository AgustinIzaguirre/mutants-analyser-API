package domain

type Dao interface {
	GetStats() (Stats, error)
}
