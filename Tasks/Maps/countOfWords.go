package Maps

import "strings"

func CountWords(sentence string) map[string]int {
	words := strings.Fields(sentence)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}
