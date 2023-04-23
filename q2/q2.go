package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {

	words := strings.Split(text, " ")

	totalLetter := 0

	for _, word := range words {
		totalLetter += len(word)

	}

	average := 0.0

	if text != "" || len(text) != 0 {
		average = float64(totalLetter) / float64(len(words))
	} else {
		return 0, fmt.Errorf("o texto está vazio")
	}

	return average, nil
}
