package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
const QUANTITY = 4

func TestMutantWithOverlappingHorizontal(t *testing.T) {
	// set up
	analyser := NewAnalyser(true, QUANTITY)
	dna := [1] string { "TTTTTA" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, true, isMutant)
}

func TestMutantWithOverlappingVertical(t *testing.T) {
	// set up
	analyser := NewAnalyser(true, QUANTITY)
	dna := [5] string { "TCA",
						"CTA",
						"CTA",
						"GCA",
						"AAA" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, true, isMutant)
}

func TestMutantWithOverlappingRightDiagonal(t *testing.T) {
	// set up
	analyser := NewAnalyser(true, QUANTITY)
	dna := [5] string { "TCAGC",
						"CTTAC",
						"CCTGG",
						"GCATC",
						"AAAGT" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, true, isMutant)
}

func TestMutantWithOverlappingLeftDiagonal(t *testing.T) {
	// set up
	analyser := NewAnalyser(true, QUANTITY)
	dna := [5] string { "TCAGC",
						"CTACC",
						"CCCGG",
						"GCATC",
						"CAAGT" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, true, isMutant)
}

func TestHumanWithoutOverlappingHorizontal(t *testing.T) {
	// set up
	analyser := NewAnalyser(false, QUANTITY)
	dna := [1] string { "TTTTTA" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, false, isMutant)
}

func TestHumanWithOverlappingVertical(t *testing.T) {
	// set up
	analyser := NewAnalyser(false, QUANTITY)
	dna := [5] string { "TCA",
						"CTA",
						"CTA",
						"GCA",
						"AAA" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, false, isMutant)
}

func TestHumanWithOverlappingRightDiagonal(t *testing.T) {
	// set up
	analyser := NewAnalyser(false, QUANTITY)
	dna := [5] string { "TCAGC",
						"CTTAC",
						"CCTGG",
						"GCATC",
						"AAAGT" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, false, isMutant)
}

func TestHumanWithOverlappingLeftDiagonal(t *testing.T) {
	// set up
	analyser := NewAnalyser(false, QUANTITY)
	dna := [5] string { "TCAGC",
						"CTACC",
						"CCCGG",
						"GCATC",
						"CAAGT" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, false, isMutant)
}

func TestMutantWithHorizontalAndVertical(t *testing.T) {
	// set up
	analyser := NewAnalyser(false, QUANTITY)
	dna := [5] string { "CCAGC",
						"CTATC",
						"CCGGG",
						"CCATC",
						"GAAAA" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, true, isMutant)
}

func TestMutantWithLeftDiagonalAndRightDiagonal(t *testing.T) {
	// set up
	analyser := NewAnalyser(false, QUANTITY)
	dna := [5] string { "CAAGG",
						"CTAGC",
						"TCGAG",
						"AGATA",
						"TAACG" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, true, isMutant)
}

func TestMutantOverlappingWithDifferentDirection(t *testing.T) {
	// set up
	analyser := NewAnalyser(false, QUANTITY)
	dna := [5] string { "GAAGG",
						"CGAGC",
						"TCGCG",
						"AGAGT",
						"ATCCA" }

	// actions
	isMutant := analyser.IsMutant(dna[:])

	// post conditions
	assert.Equal(t, true, isMutant)
}