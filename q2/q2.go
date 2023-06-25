package q2

import "strings"

func AverageLettersPerWord(text string) (float64, error) {
	words := strings.Fields(text)
	numWords := len(words)

	if numWords == 0 {
		return 0, errors.New("texto vazio")
	}

	totalLetters := 0
	for _, word := range words {
		totalLetters += len(word)
	}

	average := float64(totalLetters) / float64(numWords)
	return average, nil

}
