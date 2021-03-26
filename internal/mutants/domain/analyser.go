package domain

const (
	HORIZONTAL uint8 = 1
	VERTICAL uint8 = 2
	RIGHT_DIAGONAL uint8 = 4
	LEFT_DIAGONAL uint8 = 8
)

const DEFAULT_QUANTITY int = 4

type analyser struct {
	allowOverlapping bool
	quantity int
}

func NewAnalyser(allowOverlapping bool) *analyser {
	return &analyser{allowOverlapping: allowOverlapping, quantity: DEFAULT_QUANTITY}
}

func (analyser *analyser) IsMutant(dna[] string) bool {
	directions := [4][2] int {{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	foundChains := getFoundChainsTable(dna)
	count := 0
	for i := 0; i < len(dna); i++ {
		for j := 0; j < len(dna[i]); j++ {
			for _, direction := range directions {
				if analyser.searchInDirection(direction, i, j, dna, foundChains) {
					if !analyser.allowOverlapping {
						markAsFound(direction, i, j, foundChains, analyser.quantity);
					}
					count ++
				}
			}
			if count > 1 {
				return true
			}
		}
	}
	return count > 1
}

func (analyser *analyser) searchInDirection(direction [2] int, row, col int, dna[] string, foundChains[][] uint8) bool {
	if !analyser.allowOverlapping && isVisitedOnDirection(foundChains, row, col, direction) {
		return false
	}
	currentNucleotide := dna[row][col]
	for i := 1; i < analyser.quantity; i++ {
		row += direction[0]
		col += direction[1]
		if !isInside(row, col, dna) || currentNucleotide != dna[row][col] ||
			(!analyser.allowOverlapping && isVisitedOnDirection(foundChains, row, col, direction)) {
			return false
		}
	}
	return true
}

func isInside(row, col int, dna[] string) bool {
	if row < 0 || row >= len(dna) || col < 0 || col >= len(dna[row]) {
		return false
	}
	return true
}

func getFoundChainsTable(dna [] string) [][]uint8 {
	foundChains := make([][]uint8, len(dna))
	for i := 0; i < len(dna); i++ {
		foundChains[i] = make([]uint8, len(dna[i]))
	}
	return foundChains
}

func markAsFound(direction[2] int, row, col int, foundChains[][] uint8, quantity int) {
	directionValue := getDirectionValue(direction);
	foundChains[row][col] = foundChains[row][col] | directionValue
	for i := 0; i < quantity; i++ {
		foundChains[row][col] = foundChains[row][col] | directionValue
		row += direction[0]
		col += direction[1]
	}
}

func getDirectionValue(direction[2] int) uint8{
	if direction[0] == 0 && direction[1] == 1 {
		return HORIZONTAL
	} else if direction[0] == 1 && direction[1] == 0 {
		return VERTICAL
	} else if direction[0] == 1 && direction[1] == 1 {
		return RIGHT_DIAGONAL
	} else {
		return LEFT_DIAGONAL
	}
}

func isVisitedOnDirection(foundChains[][] uint8, row, col int, direction[2] int) bool {
	directionValue := getDirectionValue(direction)
	return (foundChains[row][col] & directionValue) > 0
}