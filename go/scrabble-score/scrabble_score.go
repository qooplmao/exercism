package scrabble

import (
	"strings"
)

// Calculate scrabble score for a given word
func Score(word string) int {
	score := 0

	for i := 0; i < len(word); i++ {
		score += letterScore(word[i : i+1])
	}

	return score
}

// Get score for a single letter, default to 0
func letterScore(letter string) int {
	switch strings.ToUpper(letter) {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	default:
		return 0
	}
}
