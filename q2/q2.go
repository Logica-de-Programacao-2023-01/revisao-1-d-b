package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {

	words := strings.Fields(text)

	if len(words) == 0 {
		return 0, fmt.Errorf("texto vazio")
	} else if len(strings.TrimSpace(text)) == 0 {
		return 0, fmt.Errorf("texto vazio")
	} else if len(text) == 0 {
		return 0, fmt.Errorf("texto vazio")
	}

	totalLetter := 0

	for _, word := range words {
		totalLetter += len(word)
	}

	average := float64(totalLetter) / float64(len(words))

	return average, nil
}
