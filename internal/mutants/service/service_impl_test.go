package service

import (
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockDao struct {
	mock.Mock
	count int
}

func (mockDao *MockDao) AddAnalysis(dna string, isMutant bool) (bool, errors.ApiError) {
	return true, nil
}

func (mockDao *MockDao) HasDNASequence(dna string) (bool, errors.ApiError) {
	if mockDao.count == 0 {
		mockDao.count++
		return false, nil
	} else {
		return true, nil
	}
}

func TestAddAnalysisWithHuman(t *testing.T) {
	// set up
	mockDao := new(MockDao)
	service := New(mockDao)
	sequence := [2] string { "TTGTCA", "CGGTGC" }
	dna := domain.New(sequence[:])

	// actions
	result, err := service.AddAnalysis(dna, false)

	// post conditions
	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func TestAddAnalysisWithMutant(t *testing.T) {
	// set up
	mockDao := new(MockDao)
	service := New(mockDao)
	sequence := [2] string { "AAAA", "GGGG" }
	dna := domain.New(sequence[:])

	// actions
	result, err := service.AddAnalysis(dna, false)

	// post conditions
	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func TestAddAnalysisWithSubmittedDNA(t *testing.T) {
	// set up
	mockDao := new(MockDao)
	service := New(mockDao)
	sequence := [2] string { "AAAA", "GGGG" }
	dna := domain.New(sequence[:])

	// actions
	first_result, err := service.AddAnalysis(dna, false)
	second_result, err2 := service.AddAnalysis(dna, false)

	// post conditions
	assert.Equal(t, true, first_result)
	assert.Nil(t, err)
	assert.Equal(t, false, second_result)
	assert.NotNil(t, err2)
}
