package domain

type Dna struct {
	sequence [] string
	validLetters map [string] bool
}

func New(sequence [] string) *Dna {
	validLetters := make(map[string]bool)
	validLetters["A"] = true
	validLetters["C"] = true
	validLetters["G"] = true
	validLetters["T"] = true

	return &Dna{sequence: sequence, validLetters: validLetters}
}

func (dna *Dna) IsValid() bool {
	if len(dna.sequence) == 0 {
		return false
	}
	for i := 0; i < len(dna.sequence); i++ {
		if len(dna.sequence[i]) == 0 {
			return false
		}
		for j := 0; j < len(dna.sequence[i]); j++ {
			if !dna.validLetters[string(dna.sequence[i][j])] {
				return false
			}
		}
	}
	return true
}

func (dna *Dna) ToString() string {
	result := ""
	for i := 0; i < len(dna.sequence); i++ {
		for j := 0; j < len(dna.sequence[i]); j++ {
			result += string(dna.sequence[i][j])
		}
		result += "_"
	}
	return result
}


func (dna *Dna) GetSequence() []string {
	return dna.sequence
}