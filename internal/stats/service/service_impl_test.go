package service

import (
"github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"
"github.com/AgustinIzaguirre/mutants-analyser-api/internal/stats/domain"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/mock"
"testing"
)

type MockDao struct {
	mock.Mock
	stats domain.Stats
}

func (mockDao *MockDao) GetStats() (domain.Stats, errors.ApiError) {
	stats := domain.Stats{1, 10, 0.1}
	return stats, nil
}

func TestAddAnalysisWithHuman(t *testing.T) {
	// set up
	mockDao := new(MockDao)
	service := New(mockDao)
	stats := domain.Stats{1, 10, 0.1}

	// actions
	result_stats, err := service.GetStats()

	// post conditions
	assert.Equal(t, stats.Mutants, result_stats.Mutants)
	assert.Equal(t, stats.Humans, result_stats.Humans)
	assert.Equal(t, stats.Ratio, result_stats.Ratio)
	assert.Nil(t, err)
}
