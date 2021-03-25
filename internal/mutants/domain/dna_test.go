package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidWithValidSequence(t *testing.T) {
	// set up
	sequence := [2] string { "TTGTCA",
							 "CGGTGC" }
	dna := New(sequence[:])

	// actions
	isValid := dna.IsValid()

	// post conditions
	assert.Equal(t, true, isValid)
}

func TestIsValidWithInvalidCharsInSequence(t *testing.T) {
	// set up
	sequence := [2] string { "TTGTCZ",
							 "CGGTGC" }
	dna := New(sequence[:])

	// actions
	isValid := dna.IsValid()

	// post conditions
	assert.Equal(t, false, isValid)
}

func TestIsValidWithLowercaseCharsInSequence(t *testing.T) {
	// set up
	sequence := [1] string { "TTGTCa" }
	dna := New(sequence[:])

	// actions
	isValid := dna.IsValid()

	// post conditions
	assert.Equal(t, false, isValid)
}

func TestIsValidWithSymbolsInSequence(t *testing.T) {
	// set up
	sequence := [2] string { "TTGTC+",
							 "AC-*A" }
	dna := New(sequence[:])

	// actions
	isValid := dna.IsValid()

	// post conditions
	assert.Equal(t, false, isValid)
}

func TestToStringOneLineSequence(t *testing.T) {
	// set up
	sequence := [1] string { "TTGTCA" }
	dna := New(sequence[:])
	expected := "TTGTCA\n"

	// actions
	actual := dna.ToString()

	// post conditions
	assert.Equal(t, expected, actual)
}

func TestToStringMultipleLineaSequence(t *testing.T) {
	// set up
	sequence := [3] string { "TTGTCA",
		"CGGTGC",
		"AGACGG" }
	dna := New(sequence[:])
	expected := "TTGTCA\nCGGTGC\nAGACGG\n"

	// actions
	actual := dna.ToString()

	// post condition
	assert.Equal(t, expected, actual)
}