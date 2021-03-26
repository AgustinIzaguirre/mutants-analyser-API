package domain

type Service interface {
	GetStats() (Stats, error)
}
