package q2

func AverageLettersPerWord(text string) (float64, error) {

	words := strings.Split(text, " ")

	totalletter := 0

	for _, word := range words {
		totalletter += len(word)

	}

	average := float64(totalletter) / float64(len(words))

	return average, nil
}

